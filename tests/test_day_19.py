import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = 'input/day_19_example.txt'
ACTUAL_INPUT = 'input/day_19.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(19, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_official_example(self, solution):
        assert solution.run_file(OFFICIAL_EXAMPLE) == '79'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '362'


@pytest.mark.parametrize('solution', [
    GoSolution(19, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_file(OFFICIAL_EXAMPLE) == '3621'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '12204'