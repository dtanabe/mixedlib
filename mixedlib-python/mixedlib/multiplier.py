from typing import Generic, TypeVar

from .operations import BinaryOperation

__all__ = ['Multiplier']

T = TypeVar('T')

class Multiplier(BinaryOperation[T]):
    def do(self, a: T, b: T) -> T:
        """
        Return the product of the two numbers.

        @param a The first number to multiply.
        @param b The second number to multiply.
        @return The product of those two numbers.
        """
        return a * b
