package cah

import "fmt"
import "math/rand"

var wcList = []string{"A vagina that leads to another dimension.", "A cat video so cute that your eyes roll back and your spine slides out of your anus.", "Going around punching people.", "The systematic destruction of an entire people and their way of life.", "Filling every orifice with butterscotch pudding.", "Slapping a racist old lady.", "Actually getting shot, for real.", "My manservant, Claude.", "Girls that always be textin'.", "A surprising amount of hair.", "Eating Tom Selleck's mustache to gain his powers.", "Samuel L. Jackson.", "An ass disaster.", "Indescribable loneliness.", "Blood farts.", "Unlimited soup, salad, and breadsticks.", "Mufasa's death scene.", "Chugging a lava lamp.", "A cop who is also a dog.", "Spending lots of money.", "A spontaneous conga line.", "The land of chocolate.", "Fisting.", "Cock.", "Crying into the pages of Sylvia Plath.", "Self-flagellation.", "The moist, demanding chasm of his mouth.", "The Harlem Globetrotters.", "Putting an entire peanut butter and jelly sandwich into the VCR.",
	"Letting everyone down.", "Not contributing to society in any meaningful way.", "A black male in his early 20s, last seen wearing a hoodie.", "All my friends dying.", "Warm, velvety muppet sex.", "Dying alone and in pain.", "Bill Clinton, naked on a bearskin rug with a saxophone.", "A boo-boo.", "A greased-up Matthew McConaughey.", "Drinking ten 5-hour ENERGY(R)s to get 50 continuous hours of energy.", "An all-midget production of Shakespeare's Richard III.", "Some douche with an acoustic guitar.", "Disco fever.", "Screaming like a maniac.", "Flying robots that kill people.", "A botched circumcision.", "Jumping out at people.", "That ass.", "Demonic possession.", "Vomiting mid-blowjob.", "Sneezing, farting, and coming at the same time.", "Three months in the hole.", "Having sex on top of a pizza.", "A pile of squirming bodies.", "Blowing some dudes in an alley.", "Getting your dick stuck in a Chinese finger trap with another dick.", "The entire Internet.", "The primal, ball-slapping sex your parents are having right now.",
	"Buying the right pants to be cool.", "The thin veneer of situational causality that underlies porn.", "Velcro(TM).", "Reverse cowgirl.", "The Quesadilla Explosion Salad(TM) from Chili.s(R).", "Nothing.", "An unstoppable wave of fire ants.", "Having shotguns for legs.", "Shutting the fuck up.", "The way white people is.", "A PowerPoint presentation.", "Roland the Farter, flatulist to the king.", "Vietnam flashbacks.", "A lamprey swimming up the toilet and latching onto your taint.", "Some kind of bird man.", "Running naked through a mall, pissing and shitting everywhere.", "Gay aliens.", "Not having sex.", "A bucket of fish heads.", "Silence.", "A sausage festival.", "An honest cop with nothing left to lose.", "Famine.", "Flesh-eating bacteria.", "Flying sex snakes.", "Not giving a shit about the Third World.", "Sexting.", "Shapeshifters.", "Porn stars.", "Raping and pillaging.", "72 virgins.", "A live studio audience.", "A time travel paradox.", "Authentic Mexican cuisine.", "Bling.", "Synergistic management solutions.", "Crippling debt.", "Daddy issues.", "Used panties.", "Dropping a chandelier on your enemies and riding the rope up.", "Former President George W. Bush.", "Full frontal nudity.", "Hormone injections.", "Laying an egg.", "Getting naked and watching Nickelodeon.", "Pretending to care.", "Public ridicule.", "Sharing needles.", "Boogers.", "The inevitable heat death of the universe.", "The miracle of childbirth.", "The Rapture.", "Whipping it out.", "White privilege.", "Wifely duties.", "The Hamburglar.", "AXE Body Spray.", "The Blood of Christ.", "Horrifying laser hair removal accidents.", "BATMAN!!!", "Agriculture.", "A robust mongoloid.", "Natural selection.", "Coat hanger abortions.", "Eating all of the cookies before the AIDS bake-sale.", "Michelle Obama's arms.", "The World of Warcraft.", "Swooping.", "Obesity.", "A homoerotic volleyball montage.", "Lockjaw.", "A mating display.", "Testicular torsion.", "All-you-can-eat shrimp for $4.99.", "Domino's(TM) Oreo(TM) Dessert Pizza.", "Kanye West.", "Hot cheese.",
	"Raptor attacks.", "Taking off your shirt.", "Smegma.", "Alcoholism.", "A middle-aged man on roller skates.", "The Care Bear Stare.", "Bingeing and purging.", "Oversized lollipops.", "Self-loathing.", "Children on leashes.", "Half-assed foreplay.", "The Holy Bible.", "German dungeon porn.", "Being on fire.", "Teenage pregnancy.", "Gandhi.", "Leaving an awkward voicemail.", "An uppercut.", "A pyramid of severed heads.", "An erection that lasts longer than four hours.", "My genitals.", "Picking up girls at the abortion clinic.", "Science.", "Not reciprocating oral sex.", "Flightless birds.", "A good sniff.", "50,000 volts straight to the nipples.", "A balanced breakfast.", "Historically black colleges.", "Actually taking candy from a baby.", "The Make-A-Wish(R) Foundation.", "A tribe of warrior women.", "Passive-aggressive Post-it notes.", "The Chinese gymnastics team.", "Switching to Geico(R).", "Peeing a little bit.", "Home video of Oprah sobbing into a Lean Cuisine(R).", "Wet dreams.", "The Jews.", "My humps.",
	"Powerful thighs.", "Winking at old people.", "Mr. Clean, right behind you.", "A caress of the inner thigh.", "Sexual tension.", "An M16 assault rifle.", "Skeletor.", "Fancy Feast(R).", "Being rich.", "Sweet, sweet vengeance.", "Republicans.", "A gassy antelope.", "Natalie Portman.", "Copping a feel.", "Kamikaze pilots.", "Sean Connery.", "The homosexual lifestyle.", "The hardworking Mexican.", "A falcon with a cap on its head.", "Altar boys.", "The Kool-Aid Man.", "Getting so angry that you pop a boner.", "Free samples.", "A big hoopla about nothing.", "Doing the right thing.", "The Three-Fifths compromise.", "Lactation.", "World peace.", "RoboCop.", "Chutzpah.", "Justin Bieber.", "Oompa-Loompas.", "Inappropriate yodeling.", "Puberty.", "Ghosts.", "An asymmetric boob job.", "Vigorous jazz hands.", "Fingering.", "Rush Limbaugh's soft, shitty body.", "GoGurt(R).", "Police brutality.", "John Wilkes Booth.", "Preteens.", "Scalping.", "Helplessly giggling at the mention of Hutus and Tutsis.", "Electricity.", "Darth Vader.",
	"A sad handjob.", "Exactly what you'd expect.", "Expecting a burp and vomiting on the floor.", "Adderall(TM).", "Embryonic stem cells.", "Tasteful sideboob.", "Panda sex.", "An icepick lobotomy.", "Tom Cruise.", "Mouth herpes.", "Sperm whales.", "Homeless people.", "Third base.", "Incest.", "Pac-Man uncontrollably guzzling cum.", "A mime having a stroke.", "Hulk Hogan.", "God.", "Scrubbing under the folds.", "Golden showers.", "Emotions.", "Licking things to claim them as your own.", "Pabst Blue Ribbon.", "The placenta.", "Spontaneous human combustion.", "Friends with benefits.", "Finger painting.", "Old-people smell.", "Dying of dysentery.", "My inner demons.", "A Super Soaker(TM) full of cat pee.", "Aaron Burr.", "Cuddling.", "The chronic.", "Battlefield amputations.", "Friendly fire.", "Ronald Reagan.", "A disappointing birthday party.", "A sassy black woman.", "A spastic nerd.", "A tiny horse.", "William Shatner.", "Riding off into the sunset.", "An M. Night Shyamalan plot twist.", "Brown people.", "Mutually-assured destruction.",
	"Pedophiles.", "Yeast.", "Grave robbing.", "Eating the last known bison.", "Catapults.", "Poor people.", "Destroying the evidence.", "The Hustle.", "The Force.", "Wiping her butt.", "Getting married, having a few kids, buying some stuff, retiring to Florida, and dying.", "Loose lips.", "AIDS.", "Pictures of boobs.", "The Ubermensch.", "Sarah Palin.", "Hospice care.", "Getting really high.", "Scientology.", "Penis envy.", "Praying the gay away.", "Frolicking.", "Two midgets shitting into a bucket.", "The KKK.", "Genghis Khan.", "Crystal meth.", "Serfdom.", "Stranger danger.", "A Bop It(TM).", "Shaquille O'Neal's acting career.", "Prancing.", "Vigilante justice.", "Overcompensation.", "Pixelated bukkake.", "A lifetime of sadness.", "Racism.", "Menstrual rage.", "Sunshine and rainbows.", "A monkey smoking a cigar.", "Rehab.", "Lance Armstrong's missing testicle.", "Dry heaving.", "The terrorists.", "Britney Spears at 55.", "Attitude.", "Breaking out into song and dance.", "Leprosy.", "Gloryholes.", "Nipple blades.",
	"The heart of a child.", "Puppies!", "Waking up half-naked in a Denny's parking lot.", "Dental dams.", "Toni Morrison's vagina.", "The taint; the grundle; the fleshy fun-bridge.", "Active listening.", "Ethnic cleansing.", "The Little Engine That Could.", "The invisible hand.", "Waiting 'til marriage.", "Unfathomable stupidity.", "Shiny objects.", "The Devil himself.", "Autocannibalism.", "Erectile dysfunction.", "My collection of high-tech sex toys.", "The Pope.", "White people.", "Tentacle porn.", "Capturing Newt Gingrich and forcing him to dance in a monkey suit.", "The penny whistle solo from 'My Heart Will Go On.'", "Seppuku.", "Same-sex ice dancing.", "Cheating in the Special Olympics.", "Throwing a virgin into a volcano.", "Keanu Reeves.", "Sean Penn.", "Nickelback.", "A look-see.", "Pooping back and forth. Forever.", "A subscription to Men's Fitness.", "Kids with ass cancer.", "A salty surprise.", "The South.", "The violation of our most basic human rights.", "YOU MUST CONSTRUCT ADDITIONAL PYLONS.", "Date rape.",
	"Being fabulous.", "Necrophilia.", "Centaurs.", "Bill Nye the Science Guy.", "Black people.", "The Boy Scouts of America.", "Lunchables(TM).", "Bitches.", "The profoundly handicapped.", "Heartwarming orphans.", "MechaHitler.", "Fiery poops.", "Another goddamn vampire movie.", "Tangled Slinkys.", "The true meaning of Christmas.", "Estrogen.", "A zesty breakfast burrito.", "That thing that electrocutes your abs.", "Passing a kidney stone.", "A bleached asshole.", "Michael Jackson.", "Cybernetic enhancements.", "Guys who don't call.", "Smallpox blankets.", "Masturbation.", "Classist undertones.", "Queefing.", "Concealing a boner.", "Edible underpants.", "Viagra(R).", "Soup that is too hot.", "Muhammad (Praise Be Unto Him).", "Surprise sex!", "Five-Dollar Footlongs(TM).", "Drinking alone.", "Dick fingers.", "Multiple stab wounds.", "Poopy diapers.", "Child abuse.", "Anal beads.", "Civilian casualties.", "Pulling out.", "Robert Downey, Jr.", "Horse meat.", "A really cool hat.", "Stalin.", "A stray pube.", "Jewish fraternities.",
	"The token minority.", "Doin' it in the butt.", "Feeding Rosie O'Donnell.", "Teaching a robot to love.", "A can of whoop-ass.", "A windmill full of corpses.", "Count Chocula.", "Wearing underwear inside-out to avoid doing laundry.", "A death ray.", "The glass ceiling.", "A cooler full of organs.", "The American Dream.", "Not wearing pants.", "When you fart and a little bit comes out.", "Take-backsies.", "Dead babies.", "Foreskin.", "Saxophone solos.", "Italians.", "A fetus.", "Firing a rifle into the air while balls deep in a squealing hog.", "Dick Cheney.", "Amputees.", "Eugenics.", "My relationship status.", "Christopher Walken.", "Bees?", "Harry Potter erotica.", "Stormtroopers.", "Getting drunk on mouthwash.", "Nazis.", "8 oz. of sweet Mexican black-tar heroin.", "Stephen Hawking talking dirty.", "Dead parents.", "Object permanence.", "Opposable thumbs.", "Racially-biased SAT questions.", "Jibber-jabber.", "Chainsaws for hands.", "Nicolas Cage.", "Child beauty pageants.", "Explosions.", "Sniffing glue.",
	"Filling Sean Hannity with helium and watching him float away.", "Repression.", "Roofies.", "My vagina.", "Assless chaps.", "A murder most foul.", "Giving 110%.", "Her Majesty, Queen Elizabeth II.", "The Trail of Tears.", "Being marginalized.", "Goblins.", "Hope.", "The Rev. Dr. Martin Luther King, Jr.", "A micropenis.", "My soul.", "A hot mess.", "Vikings.", "Hot people.", "Seduction.", "An Oedipus complex.", "Geese.", "Extremely tight pants.", "New Age music.", "Hot Pockets(R).", "Making a pouty face.", "Vehicular manslaughter.", "Women's suffrage.", "A defective condom.", "Judge Judy.", "African children.", "The Virginia Tech Massacre.", "Barack Obama.", "Asians who aren't good at math.", "Elderly Japanese men.", "The female orgasm.", "Heteronormativity.", "Parting the Red Sea.", "Arnold Schwarzenegger.", "Road head.", "Spectacular abs.", "Figgy pudding.", "A mopey zoo lion.", "A bag of magic beans.", "Poor life choices.", "My sex life.", "Auschwitz.", "A snapping turtle biting the tip of your penis.",
	"A thermonuclear detonation.", "The clitoris.", "The Big Bang.", "Land mines.", "The entire Mormon Tabernacle Choir.", "A micropig wearing a tiny raincoat and booties.", "The Dance of the Sugar Plum Fairy.", "Jerking off into a pool of children's tears.", "Man meat.", "Me time.", "The Underground Railroad.", "Poorly-timed Holocaust jokes.", "A sea of troubles.", "Lumberjack fantasies.", "Morgan Freeman's voice.", "Women in yogurt commercials.", "Natural male enhancement.", "Being a motherfucking sorcerer.", "Genital piercings.", "Passable transvestites.", "Sexy pillow fights.", "Balls.", "Grandma.", "Friction.", "Chunks of dead prostitute.", "Farting and walking away.", "Being a dick to children.", "One trillion dollars.", "The Tempur-Pedic(R) Swedish Sleep System(TM).", "Dying.", "Hurricane Katrina.", "The gays.", "The folly of man.", "Men.", "The Amish.", "Pterodactyl eggs.", "A bitch slap.", "A brain tumor.", "Cards Against Humanity.", "Fear itself.", "Lady Gaga.", "The milk man.", "A foul mouth.", "A big black dick.",
	"One thousand Slim Jims.", "Enormous Scandinavian women.", "Dorito breath.", "Shaft.", "Suicidal thoughts.", "The ooze.", "Clenched butt cheeks.", "Neil Patrick Harris.", "The economy.", "Eating an albino.", "The shambling corpse of Larry King.", "Jafar.", "Just the tip.", "Jean-Claude Van Damme.", "Ominous background music.", "Beating your wives.", "Making the penises kiss.", "Sudden Poop Explosion Disease.", "Tripping balls.", "Literally eating shit.", "The hiccups.", "The Gulags.", "Zeus's sexual appetites.", "Quiche.", "Coughing into a vagina.", "A web of lies.", "Insatiable bloodlust.", "Moral ambiguity.", "Revenge fucking.", "Scrotum tickling.", "Clams.", "A plunger to the face.", "Hipsters.", "Nubile slave boys.", "Getting abducted by Peter Pan.", "Sexy Siamese twins.", "Stockholm Syndrome.", "Space muffins.", "Leveling up.", "A nuanced critique.", "The boners of the elderly.", "Overpowering your father.", "Sexual humiliation.", "The four arms of Vishnu.", "Words, words, words.", "Ryan Gosling riding in on a white horse.",
	"Historical revisionism.", "A woman scorned.", "24-hour media coverage.", "Being a dinosaur.", "George Clooney's musk.", "A crappy little hand.", "A beached whale.", "Apologizing.", "A passionate Latino lover.", "Being a busy adult with many important things to do.", "Gladiatorial combat.", "Getting in her pants, politely.", "Mad hacky-sack skills.", "Santa Claus.", "A bloody pacifier.", "Finding a skeleton.", "Fabricating statistics.", "My machete.", "Ripping into a man's chest and pulling out his still-beating heart.", "Syphilitic insanity.", "A low standard of living.", "Deflowering the princess.", "Panty raids.", "Statistically validated stereotypes.", "Medieval Times(R) Dinner & Tournament.", "Genetically engineered super-soldiers.", "Quivering jowls.", "Bosnian chicken farmers.", "Salvia.", "Carnies.", "The harsh light of day.", "Gandalf.", "A rival dojo.", "A bigger, blacker dick.", "A sad fat dragon with no friends.", "The mere concept of Applebee's(R).", "Catastrophic urethral trauma.", "Hillary Clinton's death stare.",
	"Existing.", "Mooing.", "Swiftly achieving orgasm.", "Daddy's belt.", "Double penetration.", "Weapons-grade plutonium.", "Some really fucked-up shit.", "Subduing a grizzly bear and making her your wife.", "Rising from the grave.", "The mixing of the races.", "Taking a man's eyes and balls out and putting his eyes where his balls go and then his balls in the eye holes.", "Scrotal frostbite.", "All of this blood.", "Loki, the trickster god.", "Whining like a little bitch.", "Pumping out a baby every nine months.", "Tongue.", "Finding Waldo.", "Upgrading homeless people to mobile hotspots.", "Wearing an octopus for a hat.", "An unhinged ferris wheel rolling toward the sea.", "Living in a trashcan.", "The corporations.", "A magic hippie love cloud.", "Fuck Mountain.", "Survivor's guilt.", "Me.", "Getting hilariously gang-banged by the Blue Man Group.", "Jeff Goldblum.", "Making a friend.", "A soulful rendition of 'Ol' Man River.'", "Intimacy problems.", "A sweaty, panting leather daddy.", "Spring break!", "Being awesome at sex.",
	"Dining with cardboard cutouts of the cast of 'Friends.'", "Another shot of morphine.", "Beefin' over turf.", "A squadron of moles wearing aviator goggles.", "Bullshit.", "The Google.", "Pretty Pretty Princess Dress-Up Board Game.", "The new Radiohead album.", "An army of skeletons.", "A man in yoga pants with a ponytail and feather earrings.", "Mild autism.", "Nunchuck moves.", "Whipping a disobedient slave.", "An ether-soaked rag.", "A sweet spaceship.", "A 55-gallon drum of lube.,", "Special musical guest, Cher.", "The human body.", "Boris the Soviet Love Hammer.", "The grey nutrient broth that sustains Mitt Romney.", "Tiny nipples.", "Power.", "Oncoming traffic.", "A dollop of sour cream.", "A slightly shittier parallel universe.", "My first kill.", "Graphic violence, adult language, and some sexual content.", "Fetal alcohol syndrome.", "The day the birds attacked.", "One Ring to rule them all.", "Grandpa's ashes.", "Basic human decency.", "A Burmese tiger pit.", "Death by Steven Seagal."}
var bcList = []string{"_____________. Awesome in theory, kind of a mess in practice.", "A successful job interview begins with a firm handshake and ends with _____________.", "And what did _you_ bring for show-and-tell?", "As part of his contract, Prince won't perform without _____________ in his dressing room.", "As part of his daily regimen, Anderson Cooper sets aside 15 minutes for _____________.", "Call the law offices of Goldstein and Goldstein, because no one should have to tolerate  _____________ in the workplace.", "During high school, I never really fit in until I found _____________ Club.", "Finally! A service that delivers _____________ right to your door.", "Hey baby, come back to my place and I'll show you _____________.", "I'm not like the rest of you. I'm too rich and busy for _____________.", "In the seventh circle of Hell, sinners must endure _____________ for all eternity.", "Lovin' you is easy 'cause you're _____________.", "Money can't buy me love, but it can buy me _____________.",
	"My gym teacher got fired for adding _____________ to the obstacle course.", "The blind date was going horribly until we discovered our shared interest in _____________.", "To prepare for his upcoming role, Daniel Day-Lewis immersed himself in the world of _____________.", "Turns out that _____________-Man was neither the hero we needed nor wanted.", "What left this stain on my couch?", "What's the most emo?", "Here is the church / Here is the steeple / Open the doors / And there is _____________.", "Daddy, why is mommy crying?", "During his childhood, Salvador Dali produced hundreds of paintings of _____________.", "How did I lose my virginity?", "In 1,000 years, when paper money is a distant memory, how will we pay for goods and services?", "The Smithsonian Museum of Natural History has just opened an interactive exhibit on _____________.", "TSA guidelines now prohibit _____________ on airplanes.", "It's a pity that kids these days are all getting involved with _____________.", "Major League Baseball has banned _____________ for giving players an unfair advantage.",
	"What is Batman's guilty pleasure?", "Next from J.K. Rowling: Harry Potter and the Chamber of _____________.", "I'm sorry, Professor, but I couldn't complete my homework because of _____________.", "What did I bring back from Mexico?", "_____________. Betcha can't have just one!", "What's my anti-drug?", "While the United States raced the Soviet Union to the moon, the Mexican government funneled millions of pesos into research on _____________.", "In the new Disney Channel Original Movie, Hannah Montana struggles with _____________ for the first time.", "What's my secret power?", "What's the new fad diet?", "What did Vin Diesel eat for dinner?", "When Pharaoh remained unmoved, Moses called down a Plague of _____________.", "How am I maintaining my relationship status?", "In L.A. County Jail, word is you can trade 200 cigarettes for _____________.", "After the earthquake, Sean Penn brought _____________ to the people of Haiti.", "Instead of coal, Santa now gives the bad children _____________.",
	"Life for American Indians was forever changed when the White Man introduced them to _____________.", "What's Teach for America using to inspire inner city students to succeed?", "Maybe she's born with it. Maybe it's _____________.", "In Michael Jackson's final moments, he thought about _____________.", "White people like _____________.", "Why do I hurt all over?", "A romantic, candlelit dinner would be incomplete without _____________.", "What will I bring back in time to convince people that I am a powerful wizard?", "BILLY MAYS HERE FOR _____________.", "The class field trip was completely ruined by _____________.", "What's a girl's best friend?", "Dear Abby, I'm having some trouble with _____________ and would like your advice.", "When I am President of the United States, I will create the Department of _____________.", "What are my parents hiding from me?", "What never fails to liven up the party?", "What gets better with age?", "_____________: Good to the last drop.", "I got 99 problems but _____________ ain't one.",
	"_____________. It's a trap!", "MTV's new reality show features eight washed-up celebrities living with _____________.", "What would grandma find disturbing, yet oddly charming?", "During sex, I like to think about _____________.", "What ended my last relationship?", "What's that sound?", "_____________. That's how I want to die.", "Why am I sticky?", "What's the next Happy Meal(R) toy?", "What's there a ton of in heaven?", "I do not know with what weapons World War III will be fought, but World War IV will be fought with _____________.", "What will always get you laid?", "_____________: Kid-tested, mother-approved.", "Why can't I sleep at night?", "What's that smell?", "What helps Obama unwind?", "This is the way the world ends / This is the way the world ends / Not with a bang but with _____________.", "Coming to Broadway this season, _____________: The Musical.", "But before I kill you, Mr. Bond, I must show you _____________.", "Studies show that lab rats navigate mazes 50% faster after being exposed to _____________.",
	"Next on ESPN2: The World Series of _____________.", "When I am a billionaire, I shall erect a 50-foot statue to commemorate _____________.", "War! What is it good for?", "What gives me uncontrollable gas?", "What do old people smell like?", "What am I giving up for Lent?", "Alternative medicine is now embracing the curative powers of _____________.", "What did the US airdrop to the children of Afghanistan?", "What does Dick Cheney prefer?", "What don't you want to find in your Kung Pao chicken?", "I drink to forget _____________.", "_____________. High five, bro.", "My plan for world domination begins with _____________.", "Next season on Man vs. Wild, Bear Grylls must survive in the depths of the Amazon with only _____________ and his wits.", "Science will never explain _____________.", "The CIA now interrogates enemy agents by repeatedly subjecting them to _____________.", "In Rome, there are whisperings that the Vatican has a secret room devoted to _____________.", "When all else fails, I can always masturbate to _____________.",
	"I learned the hard way that you can't cheer up a grieving friend with _____________.", "In its new tourism campaign, Detroit proudly proclaims that it has finally eliminated _____________.", "The socialist governments of Scandinavia have declared that access to _____________ is a basic human right.", "In his new self-produced album, Kanye West raps over the sounds of _____________.", "What's the gift that keeps on giving?", "When I pooped, what came out of my butt?", "In the distant future, historians will agree that _____________ marked the beginning of America's decline.", "What has been making life difficult at the nudist colony?", "And I would have gotten away with it, too, if it hadn't been for _____________.", "What brought the orgy to a grinding halt?", "During his midlife crisis, my dad got really into _____________.", "My new favorite porn star is Joey '_____________' McGee.", "Before I run for president, I must destroy all evidence of my involvement with _____________.", "This is your captain speaking. Fasten your seatbelts and prepare for _____________.",
	"In his newest and most difficult stunt, David Blaine must escape from _____________.", "The Five Stages of Grief: denial, anger, bargaining, _____________, acceptance.", "Members of New York's social elite are paying thousands of dollars just to experience _____________.", "This month's Cosmo: 'Spice up your sex life by bringing _____________ into the bedroom.'", "Little Miss Muffet \n Sat on a tuffet, \n Eating her curds and _____________.", "My country, 'tis of thee, sweet land of _____________.", "After months of debate, the Occupy Wall Street General Assembly could only agree on 'More _____________!'", "Next time on Dr. Phil: How to talk to your child about _____________.", "Only two things in life are certain: death and _____________.", "Everyone down on the ground! We don't want to hurt anyone. We're just here for _____________.", "The healing process began when I joined a support group for victims of _____________.", "The votes are in, and the new high school mascot is _____________.", "Charades was ruined for me forever when my mom had to act out _____________.",
	"Tonight on 20/20: What you don't know about _____________ could kill you."}

//for now a "Card" is just a string, maybe later we'll mark them differently
//type Card string
type Card struct {
	Text string "json:text"
}

func (c *Card) String() string {
	return c.Text
}

func string_it(cards []*Card) []string {
	list := make([]string, 0)
	for _, value := range cards {
		list = append(list, value.Text)
	}

	return list
}

type Deck struct {
	cards map[*Card]bool
	name  string
}

func NewDeck(name string, init []string) *Deck {
	deck := &Deck{cards: make(map[*Card]bool), name: name}
	if init != nil {
		for _, item := range init {
			card := &Card{item}
			deck.cards[card] = true
		}
	}

	return deck
}

func (deck *Deck) remove(card *Card) (has bool) {
	var ok bool
	has, ok = deck.cards[card]
	if !ok {
		panic("Tried to remove a card that doesn't exist!")
	}
	delete(deck.cards, card)
	return
}

func (deck *Deck) add(card *Card) (has bool) {
	has = deck.cards[card]

	deck.cards[card] = true
	return
}

//Safetly check if a card exists
func (deck *Deck) has(card *Card) (has bool) {
	has, _ = deck.cards[card]
	return
}

func TransferSome(from *Deck, to *Deck, tomove []*Card) (ret *DeckDelta, err bool) {
	ret = &DeckDelta{Player: 0, DeckTo: to.name, DeckFrom: from.name, Cards: tomove}
	//check for trouble
	for _, item := range tomove {
		if from.has(item) != true {
			err = true
			return
		}
		if to.has(item) == true {
			err = true
			return
		}
	}
	//commit change
	for _, item := range tomove {
		from.remove(item)
		to.add(item)
	}
	return
}

func TransferAll(from *Deck, to *Deck) (ret *DeckDelta) {
	for item := range from.cards {
		from.remove(item)
		to.add(item)
	}
	ret = &DeckDelta{Player: 0, DeckTo: to.name, DeckFrom: from.name, Cards: nil}
	return
}

func (deck *Deck) randomCards() (randomList []*Card) {
	//make enough room for all the cards
	randomList = make([]*Card, len(deck.cards))
	count := 0
	//fill randomList with *Card
	for item := range deck.cards {
		randomList[count] = item
		count++
	}
	//shuffle list
	for i := cap(randomList) - 1; i > 0; i-- {
		r := rand.Int() % i
		randomList[r], randomList[i] = randomList[i], randomList[r]
	}

	return
}

func testdeck() {
	list_one, list_two := []string{"Hi", "there"}, []string{"I'm", "James"}
	deck_left := NewDeck("list-one", list_one)
	deck_right := NewDeck("list-two", list_two)

	fmt.Println("Before", deck_left, deck_right)
	deck := TransferAll(deck_left, deck_right)
	fmt.Println("After", deck_left, deck_right)

	fmt.Println("DeckDelta:", deck)

	fmt.Println("Shuffle", deck_right.randomCards())
}
