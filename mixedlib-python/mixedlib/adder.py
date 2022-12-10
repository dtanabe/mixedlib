from typing import Generic, TypeVar
from .operations import BinaryOperation

__all__ = ['Adder']

T = TypeVar('T')

class Adder(BinaryOperation[T]):
    def do(self, a: int, b: int) -> int:
        """
        Return the sum of the two numbers.

        @param a The first number to add.
        @param b The second number to add.
        @return The sum of those two numbers.
        """
        return a + b
