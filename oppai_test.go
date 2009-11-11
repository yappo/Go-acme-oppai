package oppai

import (
	"testing";
)

var tested int = 1;

func check(t *testing.T, ok bool, name string) {
	if !ok {
		t.Errorf("not ok %d - %s", tested, name)
    } else {
		t.Logf("ok %d - %s", tested, name)
	}
	tested++;
}

func testMassage() (string) {
	var oppai string;
	oppai = "　　　 _ 　∩\n";
	oppai += "　　(　゜∀゜)彡　おっぱい!おっぱい!\n";
	oppai += "　　(　 ⊂彡\n";
	oppai += "　 　|　　　|　\n";
	oppai += "　 　し ⌒Ｊ\n";
	return oppai;
}

func TestMassage(t *testing.T) {

	var o Oppai;
	check(t, o.Massage().Done() == testMassage(), "Massage 1");
	check(t, o.Massage().Massage().Done() == testMassage()+testMassage(), "Massage 2");
	check(t, o.Massage().Massage().Massage().Done() == testMassage()+testMassage()+testMassage(), "Massage 3");
}

func testOppaiUp(msg string) (aa string) {
	if msg == "" { msg = "おっぱい!" }

	aa  = "　　　 _ 　∩\n";
	aa += "　　(　゜∀゜)彡　" + msg + "\n";
	aa += "　　(　 　　|　\n";
	aa += "　 　|　　　|　\n";
	aa += "　 　し ⌒Ｊ\n";

	return aa;
}
func testOppaiDown(msg string) (aa string) {
	if msg == "" { msg = "おっぱい!" }

	aa  = "　　　 _ 　\n";
	aa += "　　(　゜∀゜)　　" + msg + "\n";
	aa += "　　(　 ⊂彡\n";
	aa += "　 　|　　　|　\n";
	aa += "　 　し ⌒Ｊ\n";

	return aa;
}

func TestOppai(t *testing.T) {
	var o Oppai;
	check(t, o.Oppai("").Done() == testOppaiUp(""), "Oppai");
	check(t, o.Oppai("").Oppai("").Done() == testOppaiUp("")+testOppaiDown(""), "Oppai Oppai");
	check(t, o.Oppai("").Oppai("").Oppai("").Done() == testOppaiUp("")+testOppaiDown("")+testOppaiUp(""), "Oppai Oppai Oppai");
	check(t,
		  o.Oppai("").Oppai("").Oppai("").Massage().Oppai("").Done() == 
		  testOppaiUp("")+testOppaiDown("")+testOppaiUp("")+testMassage()+testOppaiUp(""),
		  "Oppai Oppai Oppai Massage Oppai");
	check(t, o.Oppai("おっぱい").Oppai("だよ").Done() == testOppaiUp("おっぱい")+testOppaiDown("だよ"), "Oppai Oppai (my message)");
}
