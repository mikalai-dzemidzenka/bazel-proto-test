# Build all

```bash
bazel build //...
```

# Build only proto

```bash
bazel build //protos:go_default_library
```

## Test

```bash
bazel test //...
```
Test is randomized so if it fails use to avoid caching
```bash
bazel test //... --cache_test_results=no
```
