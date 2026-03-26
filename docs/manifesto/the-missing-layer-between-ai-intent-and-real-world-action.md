Subtitle
Why agent intelligence is not enough, and why valid state must precede real action

We are entering a new phase of AI.

For years, the main question was whether models could understand, reason, and generate.
Now the question is changing.

The new question is not only whether AI can think.
It is whether AI should be allowed to act.

This is where the current conversation is still incomplete.

Most systems today focus on how to make AI more capable.
They focus on stronger reasoning, better planning, richer tool use, longer context, more autonomous workflows, and more powerful execution.

But between AI intent and real-world action, there is still a missing layer.

That missing layer is not more intelligence.
It is not a larger model.
It is not another planning stack.
It is not just a safety wrapper.

It is a governance layer.

A layer that decides whether a candidate action has earned the right to enter the world.

That is the layer the field is still missing.

The problem is not that AI cannot act

The problem is that AI is getting increasingly good at proposing actions, while the systems around it are still weak at deciding when those actions are actually admissible.

This distinction matters.

A model can generate a highly plausible next step.
A planner can produce an elegant policy.
An agent can form a coherent multi-step strategy.
A robot can identify a path.
A drone can compute a mission fragment.

None of that means the action should happen.

A candidate is not an action.

This sounds simple.
But the future of real-world AI systems may depend on taking this distinction seriously.

Because once AI leaves the domain of pure text and enters tools, workflows, devices, robots, drones, and embodied systems, the cost of confusing candidate with action rises dramatically.

The failure mode is no longer “the answer was wrong.”
The failure mode becomes “the system changed the world when it had no right to.”

That is a deeper problem.

Most systems optimize how to do, not whether it may be done

Today’s systems are mostly optimized around capability.

Can the model solve the task?
Can the agent chain the tools?
Can the planner reach the goal?
Can the system execute without obvious failure?

These are important questions.
But they are not the first question.

The first question should be:

Does the current state legitimately support this action?

Not “Can the system produce the action?”
But “Has the system earned the right to elevate this candidate into reality?”

That shift changes everything.

Because once you ask that question, you can no longer treat planning and execution as a continuous pipeline.
You need an explicit transition point.

You need a judgment.

That judgment is what current architectures still under-specify.

State must come before intent

Before a system asks what to do next, it must establish what is currently true.

This is more than ordinary context management.

It means the system must first answer:

What is the current state?
What is the active authority?
What remains synchronized and what does not?
What has already been sealed and what is still only provisional?
Is the current container still trustworthy enough to keep governing further actions?

Without this, intent is ungrounded.

And ungrounded intent, no matter how elegant, is still structurally unsafe.

This is why state must come before intent.

Not because intelligence is unimportant.
But because intelligence without a valid current state becomes speculative power.

And speculative power is exactly what should not be allowed to directly enter the world.

Candidate is not action

This is the core sentence.

A candidate is not an action.

A proposed plan is not a command.
A model output is not an admissible instruction.
A future possibility is not a present right.

A system must be able to say:

This is a plausible action.
But it is not yet a valid one.

That is the beginning of real governance.

A mature system should not move directly from generated intent to execution.
It should move from:

candidate
to admission judgment
to executable action

That middle layer is the missing layer.

Without it, stronger agents only become faster ways to turn potential mistakes into real consequences.

Admission must come before execution

The central structural move is simple:

Execution must be preceded by admission.

A system should not ask only whether an action is possible.
It should ask whether the action is currently admissible.

That means checking at least four things:

Is the current state actually valid?
Is the authority to act actually established?
Are the evidence and source conditions sufficient?
Is the current container still fit to carry the action forward?

Only then should a system elevate a candidate into a real command.

This is not bureaucracy.
It is the minimum condition for trustworthy autonomy.

Without admission, execution becomes a guess with force.

And a guess with force is not autonomy.
It is just dangerous acceleration.

Container validity matters as much as action validity

There is a second missing insight.

A system must not only judge actions.
It must also judge itself.

Not in the philosophical sense.
In the operational sense.

Can the current runtime, session, controller, or state container still be trusted to continue governing actions?

This matters because failure does not always begin with visible collapse.

Often the deeper danger begins earlier.

The system may still appear coherent.
It may still produce clean outputs.
It may still avoid obvious symptoms.
It may even appear more stable precisely because governance has suppressed visible failure.

But underneath, control margin may already be falling.
Transport may already be degrading.
The runtime may already be leaning more heavily on hidden compensation.
The container may already be losing its qualification to continue carrying trusted state.

At that point, the right question is no longer “Is the system broken?”
The right question is:

Is this still a trustworthy state container?

This is why container validity matters.

Not only action validity.
Container validity too.

Because a system that continues governing after its container has become untrustworthy is already making invalid decisions, even if it has not visibly collapsed yet.

Handoff is not logistics

This leads to another overlooked point.

Handoff is not just transfer.
It is not file movement.
It is not administrative cleanup.

A real handoff is a governance event.

A system should not wait until obvious breakdown before transitioning state.
It should hand off when continued governance is no longer justified by the remaining control margin.

That means the trigger for transition cannot remain purely symptom-based.

Visible symptoms still matter.
But they are not enough.

Because the more governance improves, the more visible symptoms may disappear before the deeper trustworthiness problem is solved.

A mature system therefore needs to ask:

Should I continue?
Should I degrade?
Should I seal?
Should I hand off?
Should I abort?

This is not runtime hygiene.
This is the beginning of exit governance.

And exit governance is part of action governance.

Execution success does not equal legitimacy

There is one more confusion that must be rejected.

Execution success does not equal legitimacy.

A drone reaching a point does not prove it should have gone there.
A robot completing a manipulation does not prove the manipulation was admissible.
An agent finishing a workflow does not prove the workflow was elevated under a valid state.

The world contains many ways to do the wrong thing successfully.

That is why execution receipt and governance legitimacy must remain separate.

A system must preserve the distinction between:

what happened
and
what was justified

Without that distinction, any system that “works” will gradually start treating success as proof of correctness.

That is one of the fastest routes to dangerous autonomy.

Reality must re-enter after action

Even after an action is admitted and executed, the governance problem is not over.

Because action changes reality.

And once reality changes, the state used to justify the previous action is no longer the current state.

This means real-world AI systems must be reality-fed, not merely prompt-fed.

After action, reality must re-enter.

The system must take in the resulting delta, refresh the compiled state, and re-judge from there.

Otherwise the system ends up living in a past state while continuing to act in a changed world.

This is how stale state becomes structural risk.

So the governance chain is not:

state → action

It is:

reality
→ state
→ candidate
→ admission
→ execution
→ reality re-entry

That is the full loop.

The core claim

Our core claim is simple:

Only valid state may gain the right to change the world.

Everything else follows from that.

Reality must enter.
State must be established.
Intent must remain candidate until admitted.
Execution must not outrun legitimacy.
Containers must not be treated as infinitely trustworthy.
Reality must return after action.

This is not a feature.
It is not a patch.
It is not a wrapper.

It is a missing layer.

And we believe that layer will become increasingly central as AI moves from language generation into real-world action.

— Liping Wang