import pytest

from . import GoSolution

SMALL_EXAMPLE = [
    "start-A",
    "start-b",
    "A-c",
    "A-b",
    "b-d",
    "A-end",
    "b-end",
]
MEDIUM_EXAMPLE = [
    "dc-end",
    "HN-start",
    "start-kj",
    "dc-start",
    "dc-HN",
    "LN-dc",
    "HN-end",
    "kj-sa",
    "kj-HN",
    "kj-dc",
]
LARGE_EXAMPLE = [
    "fs-end",
    "he-DX",
    "fs-he",
    "start-DX",
    "pj-DX",
    "end-zg",
    "zg-sl",
    "zg-pj",
    "pj-he",
    "RW-he",
    "fs-DX",
    "pj-RW",
    "zg-RW",
    "start-pj",
    "he-WI",
    "zg-he",
    "pj-fs",
    "start-RW",
]
ACTUAL_INPUT = 'input/day_12.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(12, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_small_example(self, solution):
        assert solution.run_lines(SMALL_EXAMPLE) == '10'

    def test_medium_example(self, solution):
        assert solution.run_lines(MEDIUM_EXAMPLE) == '19'

    def test_large_example(self, solution):
        assert solution.run_lines(LARGE_EXAMPLE) == '226'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '3450'


@pytest.mark.parametrize('solution', [
    GoSolution(12, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_small_example(self, solution):
        assert solution.run_lines(SMALL_EXAMPLE) == '36'

    def test_medium_example(self, solution):
        assert solution.run_lines(MEDIUM_EXAMPLE) == '103'

    def test_large_example(self, solution):
        assert solution.run_lines(LARGE_EXAMPLE) == '3509'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '96528'
