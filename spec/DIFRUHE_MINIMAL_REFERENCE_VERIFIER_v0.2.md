# Difruhe Minimal Reference Verifier v0.1

Status: Normative Companion Specification  
Applies to: Difruhe Constitution v0.1  

---

## 1. Purpose

This document defines the minimal verification requirements
for any implementation claiming compatibility with Difruhe v0.1.

The reference verifier:

- MUST operate offline.
- MUST require no central authority.
- MUST perform deterministic validation.
- MUST not introduce semantic extensions.
- MUST not alter normative semantics defined in the Constitution.

---

## 2. Verification Scope

The minimal verifier validates:

1) Append-only structural continuity  
2) Canonical identity via rule_fingerprint  
3) Evidence binding requirement for DONE state  
4) Structural verdict state (PASS / FAIL / BLOCKED)

The verifier does NOT validate:

- Business logic
- Actor roles
- Runtime coordination
- Governance models
- Settlement mechanisms
- Economic or domain-specific semantics

---

## 3. Deterministic Canonical Procedure

Given a structural object:

The verifier MUST:

1. Canonically concatenate structural fields
   in deterministic order.

2. Compute SHA-256 over the canonical byte sequence.

3. Compare the resulting digest with the declared rule_fingerprint.

4. Validate append-only continuity via reference linkage.

5. If a structural object asserts a DONE state,
   it MUST include an evidence binding reference.

Canonicalization MUST be fully specified and documented by the implementation.

Conformance evaluation MUST bind to the declared canonicalization rules.

Implementations MUST NOT rely on implicit defaults,
hidden normalization, or environment-dependent behavior.

---

## 4. Verdict Semantics

The verifier MUST return exactly one of:

- PASS
- FAIL
- BLOCKED

Definitions:

PASS:
All structural rules satisfied.

FAIL:
A structural violation is provable from available input.

BLOCKED:
Verification cannot be completed due to missing required references
or unavailable required evidence bindings.

If a structural violation is provable from available input,
the verdict MUST be FAIL, not BLOCKED.

No additional verdict states are permitted.

Implementations MUST NOT introduce extended verdict vocabularies.

---

## 5. Determinism Requirement

Given identical input,
all compliant implementations MUST produce identical verdicts.

No randomness.
No implicit defaults.
No hidden state.
No network dependency.
No mutable global context.

Determinism is mandatory.

---

## 6. Implementation Sovereignty

Difruhe does not mandate programming language,
serialization format, storage model, or execution environment.

Any implementation is valid if:

- It adheres strictly to this specification.
- It produces deterministic verdicts.
- It does not alter structural semantics.
- It does not redefine verdict states.

Difruhe defines structure, not execution.

---

## 7. Version Binding

This verifier specification binds exclusively to:

Difruhe Constitution v0.1

Future constitutional versions require independent verifier specifications.

Implementations MUST declare which constitutional version they target.

---

End of Specification