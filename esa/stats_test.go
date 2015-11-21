package esa
import (
	"testing"
	"reflect"
)


func TestStatsGetTeamStatus(t *testing.T) {
	type TestCase struct {
		in string
		out StatsResponse
	}

	testCase := TestCase{
		in: "../tests/stubs/stats_team_stats.json",
	}

	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	res, err := client.Stats.GetTeamStats("esa")
	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %s != %s", res, testCase.out)
	}
}