package translate

import "testing"

func TestSpiltToChunks(t *testing.T) {
	text := "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet\n Nisi anim cupidatat excepteur officia\n Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident\n Nostrud officia pariatur ut officia\n Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate\n Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod\n Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim\n Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis\nLorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet\n Nisi anim cupidatat excepteur officia\n Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident\n Nostrud officia pariatur ut officia\n Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate\n Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod\n Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim\n Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis\nLorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet\n Nisi anim cupidatat excepteur officia\n Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident\n Nostrud officia pariatur ut officia\n Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate\n Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod\n Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim\n Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis\n"
	out := []string{
		"Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet\n Nisi anim cupidatat excepteur officia",
		"Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident\n Nostrud officia pariatur ut officia",
		"Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate",
		"Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod\n Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim",
		"Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis",
		"Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet\n Nisi anim cupidatat excepteur officia",
		"Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident\n Nostrud officia pariatur ut officia",
		"Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate",
		"Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod\n Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim",
		"Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis",
		"Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet\n Nisi anim cupidatat excepteur officia",
		"Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident\n Nostrud officia pariatur ut officia",
		"Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate",
		"Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod\n Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim",
		"Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis",
	}

	chunks := splitToChunks(text, 200)

	for i, chunk := range chunks {
		if chunk != out[i] {
			t.Errorf("chunk %d is not the same as out %d\n", i, i)
			t.Logf("> %v", chunk)
			t.Logf("< %v", out[i])
		}
	}
}

func TestTranslateChunk(t *testing.T) {
	text := "this is a testing text to test my function to use google translate for free"
	t.Log(text)

	_, err := TranslateChunk(text)
	if err != nil {
		t.Error(err)
	}
}

func TestTranslate(t *testing.T){
	// from https://www.pixiv.net/novel/show.php?id=19629649 and https://www.pixiv.net/novel/show.php?id=19636005
	text := "　ラズベルト・へルマンが乙女ゲームっぽいこの世界に転生したと気がついたのは、学園に入学してしばらく経った頃だった。\n\n　どの学園にも居るだろうが、多くの女生徒に人気を集めるイケメン達がいた。\n　俺様系イケメンな生徒会長の王太子に、生真面目なインテリ系イケメンの風紀委員長、可愛い弟系イケメンのクラスメイト……女生徒たちに熱い視線を向けられている彼らに対して、ラズベルトは「どこかで見たことがあるような気がするんだよなぁ」と常々思っていた。\n　しかし、どこで見たのか思い出せず、「まぁ、みんな人気者だし、どっかで見かけたんだろうな」と結論付けていた。\n\n　そんなラズベルトが、転生したと気が付いたキッカケは、レティシア・オルティースという女性徒の姿を見たときだった。婚約者である王太子の隣に立つ彼女の姿は、今まで会ったことがないはずなのに、確かに見覚えがあった。\n（彼女は、悪役令嬢の…）\n　ラズベルトは、まずそう思った。\n（あれ？　何で僕は彼女の事を悪役令嬢って思ったんだろう？　……ああ、そうかすごい昔に彼女の姿を描いているのを見たんだ……でもそれだったら、彼女の姿が今と同じなのはおかしいよな……それに物語みたいな内容だった気もするし。僕はどこで見たんだ？）\n　と、自問自答を繰り返しているうちに、自分が乙女ゲームっぽい世界に転生していたということに辿り着いたのである。「乙女ゲームの世界だ」と断言できないのは、ラズベルトの前世の記憶が、かなり断片的且つ曖昧なものだったからだ。何で自分が乙女ゲームのことを知っているのかも分からないし、内容も曖昧だ。そのため乙女ゲームのメインといえるイケメン達を見かけても「何か見たことある？」くらいの感覚だったのだ。どう考えてもイケメン達よりも登場が少ないはずのレティシアを見て気が付いたのは謎だ。\n（ここって、多分乙女ゲームの世界だよね？）\n　疑問視が付くので、「乙女ゲームっぽい世界」なのである。\n\n\n　転生したとわかったものの、ラズベルト本人は平々凡々な所謂モブだった。なので、イベントに介入することはなく、時々遠くでイベントらしき騒ぎを目撃しても「あー、もしかして、あれはイベントなのかな？」と完全に他人事であった。\n\n　普通に学園生活を送り、適度に勉学に励みながら貴族子息たちとの交流を図り、ラズベルトは卒業を迎えた。\n\n　そして卒業パーティーで、ある出来事が起こった。\n よくある『悪役令嬢の断罪イベント』である。\n\n　友人達は気になる女生徒にダンスを申し込みに行ってしまったため、ラズベルトは暇をもて余していた。ラズベルトにはダンスを申し込みたいと思う相手はおらず、そもそもこういった人が多く騒がしい場所は苦手だった。\n　早く終わらないかなと思いながら、壁際でジュースを飲み時間を潰していた。\n\n　ふと騒がしかった会場が突然シーンッと静かになったことに気がつき、ラズベルトは周りをうかがう。\n　誰か……例えば生徒会長の挨拶でも始まるのかと思った。\n　皆が注目している先に、レティシア・オルティースの姿が見えた。そして、彼女の前には王太子アレクシスと、彼に隠れるように立つ大人しそうな女性徒の姿があった。\n　生徒会長というのは合っていたが、挨拶という雰囲気ではない。\n\n　静まり返った雰囲気のなか、アレクシスが、レティシアに向かって口を開く。\n「貴女は、王太子である私の婚約者ということを傘に着て、随分と横暴な振るまいをしていたようですね。特に私と親しいからという理由でキャロルを虐げていたと聞いています。ああ、言い訳は結構。証拠は既に揃えているので。私欲にまみれた貴女は王太子の婚約者に相応しくない。よって、私はレティシア・オルティースとの婚約を破棄をここに宣言する。それに、愛想の欠片もない貴方よりも、心優しいキャロルの方が私の婚約者に相応しい」\n　言葉遣いは丁寧だが、腕を組みレティシアに向ける視線には侮蔑の色がありありと浮かんでいた。\n 　それにアレクシスの影で怯えるように立っているキャロルという女性徒──状況的にヒロインだろうが、口元に浮かぶ小さな笑みは、レティシアに勝ったという感情が伺えた。\n（見ていて気持ちの良いものではないなぁ）\n　自分の婚約者に近付く女性を牽制するのは当たり前ではないだろうか？　キャロルに対する虐めがどのようなものだったか、ラズベルトには知るよしはないが、仮に酷い内容で、王太子の婚約者として相応しくないにしても、このような衆目の前で婚約破棄をする意味がわからない。愛想云々も婚約者の条件に関係あるのだろうか？　心の優しさだけで、将来の王妃が勤まるわけないだろうに。\n（悪役令嬢も可哀想だな）\n　物語のイベントで、衆目に晒され断罪される役割だとしてもだ。\n\n（まあ、僕には関係のないことなんだけど……）\n　卒業したら領地に引っ込む予定のラズベルトには、王太子もヒロインも悪役令嬢も、今後関わることのない人物だろう。\n　それでもこれ以上この茶番劇を見るに耐えず、ラズベルトはソッと目を反らしその場を離れたのだった。　学園を卒業し、予定通り領地に戻りのんびりとした日々を過ごしていたラズベルトは、ある日父に呼ばれて書斎に赴いた。先程まで食事を共にしていたというのに何の用事なんだと思ったが、話忘れたことでもあったのだろうと、あまり深く考えなかった。\n\n「父上、来ましたよ。あれ、母上も居たのですね」\n　書斎の扉を軽くノックして、返事を待たずにラズベルトは扉を開く。なぜか母も室内にいた。\n「ラズベルト……お前は、礼儀作法をもう一度習って来た方が良いんじゃないか？」\n「家以外ならちゃんとしてますよ」\n　父が呆れたようにため息を吐くが、ラズベルトは気にしない。実際、家の外では、それなりに礼儀作法はきちんとしているつもりだ。\n「それよりも、話って何でしょう？」\n「婚約者が決まったぞ」\n「……はあ？」\n　父親から言われた内容が理解出来ず、ラズベルトは、おもいっきり首を傾げて聞き返す。\n「婚約者が決まったと言ったんだよ。相手はレティシア・オルティース公爵令嬢」\n「誰の？」\n「お前の」\n　他に誰が居るんだ？　という表情で再度告げられる。\n「レティシア・オルティースって、あのレティシア・オルティース？」\n　自分でも阿呆みたいな質問をしているなと思ったが、案の定、父は呆れ顔で「そうだ」と答えた。\n　『レティシア』という名前なら他にも居るだろうが、オルティース公爵令嬢は一人だけだ。レティシア・オルティースで間違いないのだ。\n（なんで僕と？）\n　レティシアと王太子の婚約が正式に破棄されたことは領地に戻ってしばらくしてから、風の噂で聞いた。\n　王太子が卒業パーティーで婚約破棄を宣言したからといって、王家と公爵家の政略的な婚約のはずだ。簡単には破棄出来ないのでは？　と思っていたが、随分と簡単に破棄されたことにラズベルトは驚いた。\n　そして、婚約破棄されたレティシアの新しい婚約者として、ラズベルトが選ばれたということらしい。\n　意味がわからなかった。\n（確かにへルマン家は侯爵位だけどさ）\n　公爵家に次ぐ爵位だが、へルマン侯爵家は辺境の領地を治める地味で目立たない一族だ。全体的に出世欲がなく、必要がなければ領地に引っ込んでいるような一族なので、王宮で役職を担っている者も居ない。へルマン侯爵家と聞いても、「そういえば、そんな家紋もあったかな」程度の知名度である。伯爵、男爵でもへルマン侯爵家よりも知名度の高い一族は、その辺に沢山いる。\n　そんな辺境侯爵家のラズベルトと、公爵家の中でも高位に位置するオルティース公爵家の令嬢がどうして婚約する話になるのか。\n「あと彼女、王都では色々あっただろう？　だから、お前との婚約を期にこちらの屋敷に住むことになったから」\n「は？　何でそんな展開に……」\n　王太子との婚約破棄で王都に居づらいのは理解できるが、展開が早過ぎやしないかとラズベルトは疑問に思った。\n「オルティース公爵家の別邸は王都近くばかりだからな。王都から離れるなら、いっそ一緒に住めば良いってことになったんだ」\n「三日後には屋敷に到着するらしいわよ～」\n　混乱するラズベルトに、母親が緊張感のないのほほんとした笑顔で付け加えた。\n「三日後っ！？」\n　ラズベルトは思わず大きな声になった。\n　王都からこの領地まで馬車で五日はかかるため、レティシアは既にこちらに向かっていることを意味する。\n　ということは、この婚約話は、それよりも前に決まっていたはずである。\n「あのさ……婚約の話が決まったのっていつなの？」\n　ラズベルトは恐る恐る父親に尋ねる。\n「二週間くらい前だったかな？」\n「二週間……え？　なんで、早く教えてくれなかったの？」\n　もしかして教えられない理由でもあったのだろうか。まさか、忘れてたとかバカみたいな理由であって欲しくない。\n「いやー、お前に伝えるの忘れてたんだよなぁ」\n　語尾に(笑)が見えそうなのりだった。\n　父に悪びれもなく一番嫌な理由を告げられたラズベルトは、ガックリと膝をついた。\n「そんな大事なことを伝え忘れるなーっ！！」\n　ラズベルトの叫びが、屋敷に虚しく響いたのだった。\n\n\n\n　それから三日後、レティシア・オルティース公爵令嬢が予定通りへルマン家に到着した。"
	t.Logf("length: %d\n", len(text))

	result, err := Translate(text)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
