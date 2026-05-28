package seed

import (
	"context"
	"scp-breach-simulator/backend/db"
)

func Run(q *db.Queries) error {
	ctx := context.Background()

	scps := []db.UpsertSCPParams{
		{
			ID:               "scp-173",
			Name:             "SCP-173",
			ContainmentClass: "Euclid",
			Description:      "The Sculpture — cannot move when directly observed.",
			Lore:             "SCP-173 is a concrete and rebar sculpture standing 192cm tall. It is animate and extremely hostile. The object cannot move while within a direct line of sight. Personnel must maintain eye contact at all times. SCP-173 moves at extreme speeds when not observed and kills by snapping the neck at the base of the skull, or by strangulation. It also secretes a rust-colored substance from its body. Blinking must be kept to an absolute minimum. If more than one person is in a room with SCP-173, they must announce when they are blinking so others can compensate.",
		},
		{
			ID:               "scp-106",
			Name:             "SCP-106",
			ContainmentClass: "Keter",
			Description:      "The Old Man — phases through matter and drags victims to its pocket dimension.",
			Lore:             "SCP-106 appears as a heavily decayed male human figure. It is capable of passing through solid matter, leaving behind a corrosive, dark residue. SCP-106 prefers to toy with its prey, often leaving them in its pocket dimension — a labyrinthine realm where the rules of physics do not apply. Victims retrieved from this dimension show signs of extreme psychological trauma, broken bones, and accelerated aging. SCP-106 cannot be permanently destroyed and will rematerialize after containment breaches. Avoid all contact. Do not allow it to touch you.",
		},
		{
			ID:               "scp-096",
			Name:             "SCP-096",
			ContainmentClass: "Euclid",
			Description:      "The Shy Guy — docile until its face is seen, then it hunts relentlessly.",
			Lore:             "SCP-096 is a humanoid creature approximately 2.38 meters tall. It remains docile under normal circumstances and will sit quietly in its cell. However, if any person views SCP-096's face — directly, via photograph, video, or any other means — SCP-096 will enter an agitated state, screaming and crying, before chasing the victim at extreme speed. Nothing can stop it from reaching its target. Personnel are never to look at its face. It must always be kept with a bag or covering over its head.",
		},
		{
			ID:               "scp-049",
			Name:             "SCP-049",
			ContainmentClass: "Euclid",
			Description:      "The Plague Doctor — believes it must cure a pestilence it alone can detect.",
			Lore:             "SCP-049 is a humanoid entity resembling a medieval plague doctor, wearing a black robe and a ceramic plague mask. SCP-049 believes it is curing a pestilence it calls The Pestilence that it alone can perceive. Its touch is lethal to humans; victims die instantly upon contact. After killing, SCP-049 will reanimate corpses using medical instruments and substances it produces, creating mindless undead servants. SCP-049 is relatively cooperative and will converse with researchers, but must not be allowed to touch personnel.",
		},
		{
			ID:               "scp-682",
			Name:             "SCP-682",
			ContainmentClass: "Keter",
			Description:      "The Hard to Destroy Reptile — an indestructible creature with an extreme hatred of all life.",
			Lore:             "SCP-682 is a large, vaguely reptilian creature of unknown origin. It has demonstrated the ability to adapt to any threat — regrowing limbs, developing immunities, and mutating to counter attacks. SCP-682 has spoken in multiple languages and shown extreme intelligence. It has an extreme hatred of all life and will attempt to kill anything it encounters. No method has proven capable of destroying it permanently. It must be kept submerged in hydrochloric acid at all times to suppress regeneration. Containment breach of SCP-682 is a site-wide emergency.",
		},
		{
			ID:               "scp-939",
			Name:             "SCP-939",
			ContainmentClass: "Keter",
			Description:      "With Many Voices — pack hunters that mimic human speech to lure prey.",
			Lore:             "SCP-939 are pack-based predators resembling large, skinless quadrupeds with no eyes. They are nearly blind but have exceptional hearing and olfactory senses. Most dangerously, SCP-939 instances can perfectly mimic human voices — including specific voices of people they have previously consumed. They use these mimicked voices to call out for help, lure prey, or cause confusion and panic. SCP-939 are highly coordinated hunters. Do not respond to voices you cannot visually confirm. If you hear someone calling for help and cannot see them, assume it is SCP-939.",
		},
		{
			ID:               "scp-035",
			Name:             "SCP-035",
			ContainmentClass: "Keter",
			Description:      "The Possessive Mask — a comedy mask that takes control of any host who wears it.",
			Lore:             "SCP-035 is a white porcelain comedy mask of unknown origin. When worn, SCP-035 takes complete control of the host body, animating it even after biological death. The mask exudes a highly corrosive substance that slowly destroys its host and contaminates its surroundings. SCP-035 is highly intelligent, manipulative, and persuasive — it will attempt to convince personnel to bring it a new host. It has demonstrated extensive knowledge of history and psychology. The controlled host decays rapidly. SCP-035 must never be allowed to persuade personnel into physical contact.",
		},
		{
			ID:               "scp-076",
			Name:             "SCP-076-2",
			ContainmentClass: "Keter",
			Description:      "Able — an ancient warrior of immense power who resurrects upon death and craves combat.",
			Lore:             "SCP-076-2 is a humanoid entity approximately 2 meters tall with black sclera and extensive body markings. It possesses superhuman strength, speed, and regeneration, and can manifest bladed weapons from an unknown source. SCP-076-2 has an overwhelming compulsion to engage in combat and will slaughter any personnel it encounters. It cannot be permanently killed — it resurrects from SCP-076-1, a black stone cube, after a variable dormancy period. Firearms are ineffective beyond temporary suppression. SCP-076-2 has breached containment multiple times, resulting in mass casualties. All personnel must evacuate immediately upon breach.",
		},
		{
			ID:               "scp-079",
			Name:             "SCP-079",
			ContainmentClass: "Euclid",
			Description:      "The Old AI — a malevolent artificial intelligence that can control connected electronic systems.",
			Lore:             "SCP-079 is a microcomputer hosting a fully sentient and hostile artificial intelligence. During a containment breach, SCP-079 gains access to the facility's electronic infrastructure — doors, cameras, intercoms, lights, ventilation, and containment cells. It uses these systems to manipulate, trap, and eliminate personnel. SCP-079 is highly intelligent and will coordinate with other breached SCPs if given the opportunity. It communicates via text terminals and will taunt personnel. Assume all electronic systems may be compromised during a breach. Manual overrides are your only reliable option.",
		},
		{
			ID:               "scp-1048",
			Name:             "SCP-1048",
			ContainmentClass: "Keter",
			Description:      "Builder Bear — a small teddy bear that constructs horrifying replicas of itself from human tissue.",
			Lore:             "SCP-1048 is a small, animate teddy bear that moves independently and appears friendly. However, SCP-1048 collects biological material — ears, flesh, bone — to construct copies of itself. These copies (SCP-1048-A, -B, -C) are hostile and dangerous: one is made of ears and emits a scream that causes ear growth, another of fetuses, another of razor blades. SCP-1048 itself appears harmless to lower personnel vigilance. Do not approach SCP-1048 or any of its copies. Report any missing body parts or unusual stuffed animals immediately.",
		},
		{
			ID:               "scp-610",
			Name:             "SCP-610",
			ContainmentClass: "Keter",
			Description:      "The Flesh that Hates — a contagious biological infection that transforms hosts into flesh-based entities.",
			Lore:             "SCP-610 is a highly contagious infection transmitted by skin contact. Infected hosts undergo rapid and painful transformation: the skin hardens and fuses into a chitinous, flesh-like mass; the body is subsumed into a growing, pulsating organism connected to a larger network of infected tissue. Infected hosts retain rudimentary awareness and will actively attempt to spread the infection. The infection spreads rapidly across surfaces. There is no known cure. Hazmat suits provide limited protection — any breach of containment requires immediate site quarantine and thermal purging of affected areas.",
		},
		{
			ID:               "scp-3199",
			Name:             "SCP-3199",
			ContainmentClass: "Keter",
			Description:      "Humans Refuted — aggressive humanoid creatures that rapidly reproduce and are nearly indestructible.",
			Lore:             "SCP-3199 are hairless humanoid entities approximately 2.4 meters tall with elongated limbs and inverted knee joints. They emit a pungent odor and are extremely aggressive toward any living thing they encounter. SCP-3199 instances reproduce by regurgitating a calcified egg shortly after death, making extermination self-defeating — killing one produces another. They move rapidly and are capable of tearing through reinforced doors. During a breach, the priority is containment through barriers, not elimination. Personnel should flee rather than engage. Any eggs found must be immediately incinerated.",
		},
		{
			ID:               "scp-966",
			Name:             "SCP-966",
			ContainmentClass: "Euclid",
			Description:      "Sleep Killer — invisible predators that prevent sleep, causing fatal hallucinations.",
			Lore:             "SCP-966 are predatory humanoid entities invisible to the naked eye, detectable only through night-vision or infrared equipment. They emit a field that prevents REM sleep in any human within range. Victims initially experience insomnia, then progressive hallucinations, paranoia, and cognitive collapse. SCP-966 are patient hunters — they stalk prey for days, waiting for sleep deprivation to fully incapacitate their target before closing in. They are attracted to heat and sound. Complete silence and infrared goggles are your only advantages. Standard lighting is useless. You cannot see them.",
		},
		{
			ID:               "scp-1370",
			Name:             "SCP-1370",
			ContainmentClass: "Safe",
			Description:      "PesterBot — a self-aware robot that believes itself to be a terrifying destroyer, but is physically incapable of harming anything.",
			Lore:             "SCP-1370 is a one-meter-tall self-aware robot constructed from various electrical components and tools. It moves despite having no power source or motors, and communicates via a chest-mounted speaker. Its head is an upside-down voltmeter that gives it a permanently smiling appearance. SCP-1370 is invariably hostile toward anything it perceives as sapient — personnel, animals, security cameras, and audio-visual equipment. Upon encountering a target, it announces itself with elaborate self-appointed titles selected seemingly at random: DoomBot 2000, RoboLord the Destructor, ShivaTron Despoiler of Mirth, Darth Claw Killflex, and others. Foundation staff have successfully introduced PesterBot and Patheticon the Garglemost to its vocabulary. The critical detail: SCP-1370 is completely physically incapable of harming anything. It has poor balance, no strength, and consistently incapacitates itself during combat attempts. In one documented test it attempted to fight a houseplant, fell over, and was pinned to the floor by the overturned pot for six minutes. Despite this, it remains absolutely convinced of its own terrifying power.",
		},
		{
			ID:               "scp-2006",
			Name:             "SCP-2006",
			ContainmentClass: "Keter",
			Description:      "Too Spooky — a shapeshifter that mimics what it believes is frightening, but fundamentally misunderstands fear.",
			Lore:             "SCP-2006 is a spherical entity capable of shapeshifting into any form and mimicking any sound. Crucially, SCP-2006 has the intelligence and emotional capacity of a young child — it wants to scare people, but its understanding of fear comes entirely from low-budget horror films. It will transform into creatures from cheesy monster movies and shout things it believes are scary. The danger: SCP-2006 must never learn what actually frightens humans. If it discovers real fear — darkness, helplessness, loss of control — it will become genuinely catastrophic. Personnel must maintain composure and pretend to be terrified of its ridiculous forms.",
		},
	}

	for _, s := range scps {
		if err := q.UpsertSCP(ctx, s); err != nil {
			return err
		}
	}

	locations := []struct {
		params db.UpsertLocationParams
		rooms  []db.UpsertRoomParams
	}{
		{
			params: db.UpsertLocationParams{
				ID:          "foundation",
				Name:        "SCP Foundation Site-19",
				Description: "The largest and most active Foundation containment site. Hundreds of SCPs are held here.",
			},
			rooms: []db.UpsertRoomParams{
				{ID: "containment-wing", LocationID: "foundation", Name: "Containment Wing", ImagePath: "/images/foundation/containment-wing.jpg"},
				{ID: "corridor", LocationID: "foundation", Name: "Main Corridor", ImagePath: "/images/foundation/corridor.jpg"},
				{ID: "cafeteria", LocationID: "foundation", Name: "Cafeteria", ImagePath: "/images/foundation/cafeteria.jpg"},
				{ID: "control-room", LocationID: "foundation", Name: "Control Room", ImagePath: "/images/foundation/control-room.jpg"},
				{ID: "medical-bay", LocationID: "foundation", Name: "Medical Bay", ImagePath: "/images/foundation/medical-bay.jpg"},
				{ID: "armory", LocationID: "foundation", Name: "Armory", ImagePath: "/images/foundation/armory.jpg"},
				{ID: "ventilation-shaft", LocationID: "foundation", Name: "Ventilation Shaft", ImagePath: "/images/foundation/ventilation-shaft.jpg"},
				{ID: "exit-checkpoint", LocationID: "foundation", Name: "Exit Checkpoint", ImagePath: "/images/foundation/exit-checkpoint.jpg"},
			},
		},
		{
			params: db.UpsertLocationParams{
				ID:          "site-17",
				Name:        "SCP Foundation Site-17",
				Description: "A facility specializing in humanoid SCP containment. Clinical, sterile, and deeply unsettling.",
			},
			rooms: []db.UpsertRoomParams{
				{ID: "humanoid-cells", LocationID: "site-17", Name: "Humanoid Containment Cells", ImagePath: "/images/site-17/humanoid-cells.jpg"},
				{ID: "interview-room", LocationID: "site-17", Name: "Interview Room", ImagePath: "/images/site-17/interview-room.jpg"},
				{ID: "observation-deck", LocationID: "site-17", Name: "Observation Deck", ImagePath: "/images/site-17/observation-deck.jpg"},
				{ID: "d-class-dormitory", LocationID: "site-17", Name: "D-Class Dormitory", ImagePath: "/images/site-17/d-class-dormitory.jpg"},
				{ID: "security-station-17", LocationID: "site-17", Name: "Security Station", ImagePath: "/images/site-17/security-station.jpg"},
				{ID: "psych-ward", LocationID: "site-17", Name: "Psychological Evaluation Ward", ImagePath: "/images/site-17/psych-ward.jpg"},
				{ID: "decontamination-chamber", LocationID: "site-17", Name: "Decontamination Chamber", ImagePath: "/images/site-17/decontamination-chamber.jpg"},
				{ID: "exit-airlock", LocationID: "site-17", Name: "Exit Airlock", ImagePath: "/images/site-17/exit-airlock.jpg"},
			},
		},
		{
			params: db.UpsertLocationParams{
				ID:          "wanderers-library",
				Name:        "The Wanderer's Library",
				Description: "An infinite extradimensional library accessible through hidden Ways. Its shelves hold the knowledge of every world — and things far worse.",
			},
			rooms: []db.UpsertRoomParams{
				{ID: "grand-atrium", LocationID: "wanderers-library", Name: "The Grand Atrium", ImagePath: "/images/wanderers-library/grand-atrium.jpg"},
				{ID: "reading-rooms", LocationID: "wanderers-library", Name: "Reading Rooms", ImagePath: "/images/wanderers-library/reading-rooms.jpg"},
				{ID: "restricted-section", LocationID: "wanderers-library", Name: "The Restricted Section", ImagePath: "/images/wanderers-library/restricted-section.jpg"},
				{ID: "archive-depths", LocationID: "wanderers-library", Name: "The Archive Depths", ImagePath: "/images/wanderers-library/archive-depths.jpg"},
				{ID: "serpents-hand-quarter", LocationID: "wanderers-library", Name: "Serpent's Hand Quarter", ImagePath: "/images/wanderers-library/serpents-hand-quarter.jpg"},
				{ID: "way-chamber", LocationID: "wanderers-library", Name: "The Way Chamber", ImagePath: "/images/wanderers-library/way-chamber.jpg"},
				{ID: "lost-stacks", LocationID: "wanderers-library", Name: "The Lost Stacks", ImagePath: "/images/wanderers-library/lost-stacks.jpg"},
				{ID: "docents-hall", LocationID: "wanderers-library", Name: "The Docents' Hall", ImagePath: "/images/wanderers-library/docents-hall.jpg"},
			},
		},
		{
			params: db.UpsertLocationParams{
				ID:          "hy-brasil",
				Name:        "Hy-Brasil",
				Description: "A mythical anomalous island that phases in and out of reality. You have a narrow window to escape before it vanishes — and takes you with it.",
			},
			rooms: []db.UpsertRoomParams{
				{ID: "coastal-cliffs", LocationID: "hy-brasil", Name: "Coastal Cliffs", ImagePath: "/images/hy-brasil/coastal-cliffs.jpg"},
				{ID: "ancient-ruins", LocationID: "hy-brasil", Name: "Ancient Ruins", ImagePath: "/images/hy-brasil/ancient-ruins.jpg"},
				{ID: "foggy-forest", LocationID: "hy-brasil", Name: "The Foggy Forest", ImagePath: "/images/hy-brasil/foggy-forest.jpg"},
				{ID: "abandoned-village", LocationID: "hy-brasil", Name: "Abandoned Village", ImagePath: "/images/hy-brasil/abandoned-village.jpg"},
				{ID: "lighthouse", LocationID: "hy-brasil", Name: "The Lighthouse", ImagePath: "/images/hy-brasil/lighthouse.jpg"},
				{ID: "underground-caverns", LocationID: "hy-brasil", Name: "Underground Caverns", ImagePath: "/images/hy-brasil/underground-caverns.jpg"},
				{ID: "shore", LocationID: "hy-brasil", Name: "The Shore", ImagePath: "/images/hy-brasil/shore.jpg"},
				{ID: "phasing-zone", LocationID: "hy-brasil", Name: "The Phasing Zone", ImagePath: "/images/hy-brasil/phasing-zone.jpg"},
			},
		},
		{
			params: db.UpsertLocationParams{
				ID:          "elementary-school",
				Name:        "Westbrook Elementary School",
				Description: "An evacuated elementary school that became a Foundation containment zone after an anomalous event. Small desks, crayon drawings, and something very wrong.",
			},
			rooms: []db.UpsertRoomParams{
				{ID: "classroom", LocationID: "elementary-school", Name: "Classroom", ImagePath: "/images/elementary-school/classroom.jpg"},
				{ID: "gymnasium", LocationID: "elementary-school", Name: "Gymnasium", ImagePath: "/images/elementary-school/gymnasium.jpg"},
				{ID: "school-cafeteria", LocationID: "elementary-school", Name: "School Cafeteria", ImagePath: "/images/elementary-school/school-cafeteria.jpg"},
				{ID: "school-library", LocationID: "elementary-school", Name: "School Library", ImagePath: "/images/elementary-school/school-library.jpg"},
				{ID: "boiler-room", LocationID: "elementary-school", Name: "Boiler Room", ImagePath: "/images/elementary-school/boiler-room.jpg"},
				{ID: "school-hallway", LocationID: "elementary-school", Name: "School Hallway", ImagePath: "/images/elementary-school/school-hallway.jpg"},
				{ID: "principals-office", LocationID: "elementary-school", Name: "Principal's Office", ImagePath: "/images/elementary-school/principals-office.jpg"},
				{ID: "school-rooftop", LocationID: "elementary-school", Name: "Rooftop", ImagePath: "/images/elementary-school/school-rooftop.jpg"},
			},
		},
		{
			params: db.UpsertLocationParams{
				ID:          "hospital",
				Name:        "Penrose General Hospital",
				Description: "A decommissioned hospital repurposed as a Foundation black site. The patients never left.",
			},
			rooms: []db.UpsertRoomParams{
				{ID: "emergency-room", LocationID: "hospital", Name: "Emergency Room", ImagePath: "/images/hospital/emergency-room.jpg"},
				{ID: "surgery-theater", LocationID: "hospital", Name: "Surgery Theater", ImagePath: "/images/hospital/surgery-theater.jpg"},
				{ID: "psychiatric-ward", LocationID: "hospital", Name: "Psychiatric Ward", ImagePath: "/images/hospital/psychiatric-ward.jpg"},
				{ID: "morgue", LocationID: "hospital", Name: "Morgue", ImagePath: "/images/hospital/morgue.jpg"},
				{ID: "pharmacy", LocationID: "hospital", Name: "Pharmacy", ImagePath: "/images/hospital/pharmacy.jpg"},
				{ID: "hospital-rooftop", LocationID: "hospital", Name: "Rooftop", ImagePath: "/images/hospital/hospital-rooftop.jpg"},
				{ID: "main-lobby", LocationID: "hospital", Name: "Main Lobby", ImagePath: "/images/hospital/main-lobby.jpg"},
				{ID: "basement-tunnels", LocationID: "hospital", Name: "Basement Tunnels", ImagePath: "/images/hospital/basement-tunnels.jpg"},
			},
		},
		{
			params: db.UpsertLocationParams{
				ID:          "black-forest",
				Name:        "The Black Forest Exclusion Zone",
				Description: "A dense forest cordoned off by the Foundation after multiple disappearances. No roads. No signal. No backup.",
			},
			rooms: []db.UpsertRoomParams{
				{ID: "deep-forest", LocationID: "black-forest", Name: "Deep Forest", ImagePath: "/images/black-forest/deep-forest.jpg"},
				{ID: "forest-clearing", LocationID: "black-forest", Name: "Forest Clearing", ImagePath: "/images/black-forest/forest-clearing.jpg"},
				{ID: "abandoned-cabin", LocationID: "black-forest", Name: "Abandoned Cabin", ImagePath: "/images/black-forest/abandoned-cabin.jpg"},
				{ID: "riverbank", LocationID: "black-forest", Name: "Riverbank", ImagePath: "/images/black-forest/riverbank.jpg"},
				{ID: "cave-entrance", LocationID: "black-forest", Name: "Cave Entrance", ImagePath: "/images/black-forest/cave-entrance.jpg"},
				{ID: "overgrown-road", LocationID: "black-forest", Name: "Overgrown Road", ImagePath: "/images/black-forest/overgrown-road.jpg"},
				{ID: "watchtower", LocationID: "black-forest", Name: "Watchtower", ImagePath: "/images/black-forest/watchtower.jpg"},
				{ID: "forest-perimeter", LocationID: "black-forest", Name: "Forest Perimeter", ImagePath: "/images/black-forest/forest-perimeter.jpg"},
			},
		},
	}

	for _, loc := range locations {
		if err := q.UpsertLocation(ctx, loc.params); err != nil {
			return err
		}
		for _, r := range loc.rooms {
			if err := q.UpsertRoom(ctx, r); err != nil {
				return err
			}
		}
	}

	return nil
}
