import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = [
    "forward 5",
    "down 5",
    "forward 8",
    "up 3",
    "down 8",
    "forward 2",
]
ACTUAL_INPUT = 'input/day_02.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(2, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '150'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '1690020'


@pytest.mark.parametrize('solution', [
    GoSolution(2, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '900'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '1408487760'
