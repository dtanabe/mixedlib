from typing import Generic, TypeVar

__all__ = ['BinaryOperation']

T = TypeVar('T')

class BinaryOperation(Generic[T]):
    def do(self, a: T, b: T) -> T:
        raise NotImplementedError("BinaryOperation.do must be implemented")
