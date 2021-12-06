import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = ["3,4,3,1,2"]
ACTUAL_INPUT = 'input/day_06.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(6, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '5934'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '375482'


@pytest.mark.parametrize('solution', [
    GoSolution(6, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '26984457539'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '1689540415957'
