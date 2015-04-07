/*
 * Copyright (C) 2015 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy of
 * the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

package com.google.cloud.dataflow.sdk.transforms.windowing;

import com.google.cloud.dataflow.sdk.transforms.windowing.Trigger.AtMostOnceTrigger;
import com.google.common.base.Preconditions;

import org.joda.time.Instant;

import java.util.Arrays;
import java.util.List;

/**
 * Create a {@link CompositeTrigger} that fires once after all of its sub-triggers have fired. If
 * any of the sub-triggers finish without firing, the {@code AfterAll.of(...)} will also finish
 * without firing.
 *
 * @param <W> {@link BoundedWindow} subclass used to represent the windows used by this
 *            {@code Trigger}
 */
public class AfterAll<W extends BoundedWindow>
    extends CompositeTrigger<W> implements AtMostOnceTrigger<W> {

  private static final long serialVersionUID = 0L;

  private AfterAll(List<Trigger<W>> subTriggers) {
    super(subTriggers);
    Preconditions.checkArgument(subTriggers.size() > 1);
  }

  @SafeVarargs
  public static <W extends BoundedWindow> AtMostOnceTrigger<W> of(
      AtMostOnceTrigger<W>... triggers) {
    return new AfterAll<W>(Arrays.<Trigger<W>>asList(triggers));
  }

  private TriggerResult wrapResult(SubTriggerExecutor subExecutor) {
    // If all children have finished, then they must have each fired at least once.
    if (subExecutor.allFinished()) {
      return TriggerResult.FIRE_AND_FINISH;
    }

    return TriggerResult.CONTINUE;
  }

  @Override
  public TriggerResult onElement(TriggerContext<W> c, OnElementEvent<W> e) throws Exception {
    SubTriggerExecutor subExecutor = subExecutor(c, e.window());
    for (int i : subExecutor.getUnfinishedTriggers()) {
      // Mark any fired triggers as finished.
      if (subExecutor.onElement(c, i, e).isFire()) {
        subExecutor.markFinished(c, i);
      }
    }

    return wrapResult(subExecutor);
  }

  @Override
  public TriggerResult onMerge(TriggerContext<W> c, OnMergeEvent<W> e) throws Exception {
    SubTriggerExecutor subExecutor = subExecutor(c, e);

    // If after merging the set of fire & finished sub-triggers, we're done, we can
    // FIRE_AND_FINISH early.
    if (subExecutor.allFinished()) {
      return TriggerResult.FIRE_AND_FINISH;
    }

    // Otherwise, merge all of the unfinished triggers.
    for (int i : subExecutor.getUnfinishedTriggers()) {
      if (subExecutor.onMerge(c, i, e).isFire()) {
        subExecutor.markFinished(c, i);
      }
    }

    return wrapResult(subExecutor);
  }

  @Override
  public TriggerResult afterChildTimer(
      TriggerContext<W> c, W window, int childIdx, TriggerResult result) throws Exception {
    if (TriggerResult.CONTINUE.equals(result)) {
      return TriggerResult.CONTINUE;
    }

    return wrapResult(subExecutor(c, window));
  }

  @Override
  public boolean willNeverFinish() {
    // even if one of the triggers never finishes, the AfterAll could finish if it FIREs.
    return false;
  }

  @Override
  public Instant getWatermarkCutoff(W window) {
    // This trigger will fire after the latest of its sub-triggers.
    Instant deadline = BoundedWindow.TIMESTAMP_MIN_VALUE;
    for (Trigger<W> subTrigger : subTriggers) {
      Instant subDeadline = subTrigger.getWatermarkCutoff(window);
      if (deadline.isBefore(subDeadline)) {
        deadline = subDeadline;
      }
    }
    return deadline;
  }
}
