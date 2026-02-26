## Difruhe Release v0.2.0
Release Type: Open Constitution Pack  
Language Version: v0.2
Status: PASS  

Freeze PASS receipt  
verdict=PASS repo=difruhe ref=refs/tags/v0.2.0 head=9ae5326a1723f716aa3b2271bbe11a92de17e994  
spec_fingerprint=sha256:7720458a5e5361603ba389f10e4cdc798af6324cdc27a2fa9da12ff12a1a8d4e  
Non-normative boundary: docs/ROLE_BOUNDARY_v0.2.md  
────────────────────────

## Release Scope

This release packages the finalized Difruhe v0.2 structural language.

Difruhe is a minimal structural language for irreversible, referenceable progression.

Difruhe is not a tool.  
Difruhe is not a platform.  
Difruhe is a structural foundation.

Specification is open.  
Implementation is sovereign.  
Verification is offline-capable.

────────────────────────

## Language Version

The Difruhe language version is:

v0.2

────────────────────────

## Included Specifications (Normative)

The language identity of Difruhe v0.2 is defined exclusively by:

1. DIFRUHE_CONSTITUTION_v0.2.md  
2. EVENT_VOCAB_v0.2.md  
3. REFERENCE_CANON_v0.2.md  
4. RULE_CANON_v0.2.md  

Identity is computed via:

spec_fingerprint = sha256(canonical_concat(locked_spec_files))

No other repository files participate in identity.

────────────────────────

## Structural Principles

### Foundational Axioms

- Differences are append-only  
- Differences must be referenceable  

### Structural Laws

- Append-only time  
- Event-driven state  
- Canonical identity  
- Evidence-bound completion  
- Shared structural verdict  

### Event Vocabulary (v0.2 — Final Locked)

INIT  
DECLARE  
SNAPSHOT  
EVIDENCE_ATTACH  
CHECKPOINT  
COMMIT  
VERDICT  

No additional structural events are defined in v0.2.

────────────────────────

## Reference Canon (v0.2)

A valid reference must:

- Be resolvable  
- Be verifiable  
- Be non-drifting  

No verifiable reference → no DONE.

────────────────────────

## Rule Canon (v0.2)

Rules are:

- Versioned  
- Fingerprinted  
- Non-drifting  

Completion is always relative to a specific rule version + fingerprint.

No silent rule upgrades are permitted.

────────────────────────

## Minimal Reference Verifier

A compliant verifier MUST:

- Operate offline  
- Be deterministic  
- Require no central authority  
- Produce exactly one verdict:

PASS  
FAIL  
BLOCKED  

No additional verdict states are permitted.

────────────────────────

## SHA256 Manifest

All release artifacts are fingerprinted in:

spec/sha256_manifest.v0.2.txt

The manifest is a release record.  
It does NOT participate in identity computation.  
The manifest MUST NOT include itself.

────────────────────────

## Compatibility Statement

Implementations claiming compatibility with earlier Difruhe versions (including v0.1) MUST:

- Respect append-only structural progression  
- Preserve canonical identity  
- Bind DONE to verifiable evidence  
- Produce deterministic verdicts  

Compatibility means structural preservation, not semantic reinterpretation.

Language semantics MUST NOT drift across versions.

────────────────────────

Authored by Liping Wang  
2026