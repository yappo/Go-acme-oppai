package oppai

import "container/vector"

type Oppai struct {
	count int;
	cmd vector.Vector;
}

func (self *Oppai) Massage() *Oppai {
	self.cmd.Push(map[string] string { "name": "Massage", "message": "" } );
	return self;
}

func (self *Oppai) Oppai(msg string) *Oppai {
	self.cmd.Push(map[string] string { "name": "Oppai", "message": msg } );
	return self;
}

func (self *Oppai) Done() (buf string) {
	var c int = 0;

	for i := 0; i < self.cmd.Len(); i++ {
		var cmd string = self.cmd.At(i).(map[string] string)["name"];
		var msg string = self.cmd.At(i).(map[string] string)["message"];

		if cmd == "Massage" {
			buf += doneMassage();
			c = 0;
		} else if cmd == "Oppai" {
			if c % 2 == 0 {
				buf += doneOppaiUp(msg);
			} else {
				buf += doneOppaiDown(msg);
			}
			c++;
		}
	}

	self.cmd.Init(0);
	return buf;
}

func doneMassage() (aa string) {
	aa  = "　　　 _ 　∩\n";
	aa += "　　(　゜∀゜)彡　おっぱい!おっぱい!\n";
	aa += "　　(　 ⊂彡\n";
	aa += "　 　|　　　|　\n";
	aa += "　 　し ⌒Ｊ\n";	
	return aa;
}

func doneOppaiUp(msg string) (aa string) {
	if msg == "" { msg = "おっぱい!" }

	aa  = "　　　 _ 　∩\n";
	aa += "　　(　゜∀゜)彡　" + msg + "\n";
	aa += "　　(　 　　|　\n";
	aa += "　 　|　　　|　\n";
	aa += "　 　し ⌒Ｊ\n";

	return aa;
}
func doneOppaiDown(msg string) (aa string) {
	if msg == "" { msg = "おっぱい!" }

	aa  = "　　　 _ 　\n";
	aa += "　　(　゜∀゜)　　" + msg + "\n";
	aa += "　　(　 ⊂彡\n";
	aa += "　 　|　　　|　\n";
	aa += "　 　し ⌒Ｊ\n";

	return aa;
}

