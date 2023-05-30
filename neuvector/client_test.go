package neuvector

import (
	"testing"
)

const (
	ExpiredToken = "sivfhwlJKiBN881Bmeq9kOisr1PjonDfmIOooB1jiYoW_atbDbCM2JOqF21GaPJFd0Euoicn7e+HZfWrr80BIEnzQx0SUw7Unu8N7ysfHUaqoIQWM2B1k5BTHJHhp5cfmUijuKGKt_kc1UAktG371RRunBcbQLJLjTPJZA+wGEuQGpX0OpRiYIXclOlYlEip8QLYXKAvNmyipZACbYUcBA0JkuXSBvT0b0flaGlA6atA4R6oLtv44pVTqMasq+a1rQd0UnlMh632PI8R5pDbmJgHx2Y+3yfCA6FD0fEFq_FCi36cWwqSOs7gcqm0RQYVbDiHDdk_cPQNky+7u+nZqAVeWyLnxs+mmZobJ8N0lafEdzi_Bs5bQ28826wTXueBe+Nd4mAzEljCM22xaz5nA42ftrXXLcLdpE2L8SEzaCrFRqsy4snH9hD7wpZjDV+djCHbZKrTJzBwNF3Qibhga8bOpRPWNIzQo91iSj5fIPbvbnEQs88VDC9jk6oywaiZX1xdlJHpVz9ASE_sBWgeaIb6ufwVFphWrc2AND9NpnDROaXH1gH_JHUowZCzB99oFaAy1Pz58hXhTzmlqw_t2Xrv8fA4mAZIVwuy0_j+o3fF2ca04QM0M_MRbC3iSfX+ORQzbFVO3Eqh2sLqbN5tzXpp41s+U_ee3cDuZR1UvCEfBy9EUbJcflg9yIZvAXvGCyHpRujgTJcMb59OLC5Phqd+7WeO7gXs6dayR8YSq8ciqiEYt7tPYBHmCHyr37WQ3q_RT2HYAyQQjFvhNz98S8WAZwHi6Rng7hijyAUq_vPuwIR4xefCfKfh1vhFQ3op5bZNQaLIKDdnZ4Yii++tkymmnL02ciOR_9if0pOgi_J00eK5Z+iuLrEFqADYDxaM_0hIptKJG8+IWs9syK1KWvEPZtS5_p4y7Ak0hu01iDn0e66OjFbJ4BiOt5UqtgkKgwUYNCQV_2ccaBrVYtqbsOldWv1yOjTfO0gsIvT3tEaGmgQdDY11pBRPPV4Ww8VPu0wv8A+suFH+4dQIX_NleuldkZNxNbQVBIGAGn0lNEZn4o24JlrWlqtnmrvjhbZEZnAa5StosyrZOMT_tk8Kz0Q54ECe2eeB77fADMyvqKCLM7ft6Wf8pWUY"
)

// Test if the client is able to authenticate
func TestClient(t *testing.T) {
	_, err := NewDefaultClient()

	if err != nil {
		t.Errorf("%s", err.Error())
	}
}

func TestExpiredToken(t *testing.T) {
	var ret any

	client, err := NewDefaultClient()

	if err != nil {
		t.Errorf("%s", err.Error())
	}

	// Replaced old token without refreshing aka compute a new one
	client.SetToken(ExpiredToken)

	if err := client.Get("/policy/rule", ret); err == nil {
		t.Errorf("%s", err.Error())
	}
}
