# Difruhe Identity Law v0.2 (Final Locked)

Status: Normative Companion Specification  
Applies to: Difruhe Constitution v0.2  

---

## 1) Language Identity Scope

Difruhe is a language.

Language identity MUST be computed exclusively from normative specification documents listed below.

Identity scope is exactly these 4 files, in this locked order:

1. DIFRUHE_CONSTITUTION_v0.2.md
2. EVENT_VOCAB_v0.2.md
3. REFERENCE_CANON_v0.2.md
4. RULE_CANON_v0.2.md

No other repository files may affect language identity, including:

- README
- RELEASE_NOTES
- LICENSE
- tools
- verifier implementations
- runtime modules
- git tags
- commit messages
- release artifacts

Any modification to any file listed above changes the language identity.

---

## 2) Identity Fingerprint Definition

spec_fingerprint = sha256( canonical_concat(spec_files_in_locked_order) )

Canonicalization rules (normative and mandatory):

- Read each file as UTF-8 text.
- Normalize line endings to LF (`\n`).
- Ensure exactly one trailing LF per file.
- Concatenate as:

  "FILE:<name>\n" + <content>

for each file, joined sequentially in the locked order defined above.

No additional whitespace, encoding normalization,
or implicit transformations are permitted.

---

## 3) Non-Recursive Release Law

Tagging MUST be a pointer-only operation.

Tagging MUST NOT:

- Modify any tracked files.
- Introduce additional commits.
- Regenerate identity inputs.
- Change spec_fingerprint.

Release artifacts (notes, snapshots, manifests) are release records only.

Release artifacts MUST NOT participate in identity fingerprint computation.

sha256_manifest files MUST NOT include themselves
or any release-only artifacts in their computation scope.

Language identity is independent of release mechanics.

---

End of Specification