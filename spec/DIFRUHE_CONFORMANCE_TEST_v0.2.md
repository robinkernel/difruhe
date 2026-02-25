# Difruhe Conformance Test v0.1

Status: Normative Companion Specification  
Applies to:
- Difruhe Constitution v0.1
- Difruhe Minimal Reference Verifier v0.1

This document defines the minimal conformance tests for any implementation
claiming compatibility with Difruhe v0.1.

All tests MUST be runnable offline.

---

## 1. Conformance Rule

A claim of "Difruhe v0.1 compatible" is valid only if:

- The implementation produces deterministic results, and
- It passes ALL tests defined in this document.

No partial compliance claims are permitted.

---

## 2. Test Input Format (Minimal)

A test input MUST be representable as a single UTF-8 text file containing:

- A case identifier
- A structural object payload (format is implementation-defined)
- Expected verdict: PASS / FAIL / BLOCKED

The conformance suite does not mandate a serialization format.
However, each test MUST be unambiguous and reproducible.

---

## 3. Verdict Contract

The implementation MUST return exactly one verdict:

- PASS
- FAIL
- BLOCKED

No additional verdict states are allowed.

---

## 4. Determinism Tests

### CT-001 Determinism (Same Input, Same Verdict)

Input: Any valid structural object.
Run verification 10 times.

Expected:
- Verdict MUST be identical across all runs.

Failure Mode:
- Any mismatch => FAIL conformance.

---

### CT-002 Determinism (No Hidden Defaults)

Input: A structural object with an omitted optional field.

Expected:
- Implementation MUST NOT "invent" defaults that change the canonical bytes.
- If the missing field is required for verification, verdict MUST be BLOCKED.

Failure Mode:
- Silent acceptance with implied defaults => FAIL conformance.

---

## 5. Append-Only Continuity Tests

### CT-010 Append-Only Continuity (Valid Chain)

Input:
- A sequence of structural objects A -> B -> C
- Each object references the prior object as its structural predecessor.

Expected:
- PASS if all references resolve and continuity is verifiable.

---

### CT-011 Append-Only Violation (Broken Chain)

Input:
- A sequence A -> B -> C
- C references a non-existing predecessor or an unrelated predecessor.

Expected:
- FAIL if a structural violation is provable,
- BLOCKED if required references are missing and cannot be resolved offline.

---

## 6. Canonical Identity / Fingerprint Tests

### CT-020 Fingerprint Match (Valid)

Input:
- A structural object with declared rule_fingerprint.
- Canonical bytes must be defined by the implementation, but must be deterministic.

Expected:
- PASS if computed fingerprint equals declared rule_fingerprint.

---

### CT-021 Fingerprint Mismatch (Invalid)

Input:
- Same as CT-020, but declared rule_fingerprint is intentionally incorrect.

Expected:
- FAIL.

---

## 7. Evidence-Bound DONE Tests

### CT-030 Evidence Required for DONE

Input:
- A structural object declaring a DONE state without evidence binding reference.

Expected:
- FAIL (if DONE is asserted without evidence)
OR
- BLOCKED (if the object cannot be verified due to missing required evidence reference rules)

Constraint:
- Under no circumstances may this produce PASS.

---

### CT-031 Evidence Present for DONE

Input:
- A structural object declaring DONE state with an evidence binding reference.

Expected:
- PASS if evidence binding is present AND structurally well-formed.
- BLOCKED if evidence reference is required to resolve but is unavailable offline.
- FAIL if evidence binding is structurally invalid.

---

## 8. Shared Cadence for Verdict

### CT-040 Verdict Vocabulary Lock

Input:
- A structural object where an implementation attempts to emit a non-standard verdict,
  e.g., "WARN", "UNKNOWN", "SOFT_FAIL", etc.

Expected:
- FAIL conformance.

---

## 9. Minimal Test Set Requirement

A conformance pack MUST include at minimum:

- CT-001
- CT-002
- CT-010
- CT-011
- CT-020
- CT-021
- CT-030
- CT-031
- CT-040

---

## 10. Version Binding

This conformance specification binds exclusively to Difruhe Constitution v0.1.

Future constitutional versions require independent conformance test specifications.

---

End of Specification