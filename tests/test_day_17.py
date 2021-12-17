import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = ["target area: x=20..30, y=-10..-5"]
ACTUAL_INPUT = 'input/day_17.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(17, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '45'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '4005'


@pytest.mark.parametrize('solution', [
    GoSolution(17, 2),
], ids=lambda s: s.name)
class TestPart2(object):


    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '112'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == ''
