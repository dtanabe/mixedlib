.PHONY: test
test: python-test

.PHONY: python-test
python-test: .venv/bin/tox
	$<
	
.venv/bin/tox:
	python3 -m venv .venv
	.venv/bin/pip install --upgrade pip tox

