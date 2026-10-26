.PHONY: test-tag

# Harmless controlled target. It does not build or test Aiven code.
test-tag:
	@echo "SAFE_TEST_TARGET_REACHED tag=$(tag)"
