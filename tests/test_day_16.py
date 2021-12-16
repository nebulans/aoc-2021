import pytest

from . import GoSolution

ACTUAL_INPUT = 'input/day_16.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(16, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    @pytest.mark.parametrize("encoded, expected", [
        ("D2FE28", "6"),
        ("38006F45291200", "9"),
        ("EE00D40C823060", "14"),
        ("8A004A801A8002F478", "16"),
        ("620080001611562C8802118E34", "12"),
        ("C0015000016115A2E0802F182340", "23"),
        ("A0016C880162017C3686B18A3D4780", "31"),
    ])
    def test_official_examples(self, solution, encoded, expected):
        assert solution.run_lines([encoded]) == expected

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '889'


@pytest.mark.parametrize('solution', [
    GoSolution(16, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    @pytest.mark.parametrize("encoded, expected", [
        ("C200B40A82", "3"),
        ("04005AC33890", "54"),
        ("880086C3E88112", "7"),
        ("CE00C43D881120", "9"),
        ("D8005AC2A8F0", "1"),
        ("F600BC2D8F", "0"),
        ("9C005AC2F8F0", "0"),
        ("9C0141080250320F1802104A08", "1"),
    ])
    def test_official_examples(self, solution, encoded, expected):
        assert solution.run_lines([encoded]) == expected

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '739303923668'
