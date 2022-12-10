import mixedlib


def test_adder():
    op = mixedlib.Adder()
    assert op.do(40, 2) == 42
