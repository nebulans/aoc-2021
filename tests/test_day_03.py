import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = [
    "00100",
    "11110",
    "10110",
    "10111",
    "10101",
    "01111",
    "00111",
    "11100",
    "10000",
    "11001",
    "00010",
    "01010",
]
ACTUAL_INPUT = 'input/day_03.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(3, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '198'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '1025636'


@pytest.mark.parametrize('solution', [
    GoSolution(3, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '230'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '793873'
