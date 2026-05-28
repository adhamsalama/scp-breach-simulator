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
	}

	for _, s := range scps {
		if err := q.UpsertSCP(ctx, s); err != nil {
			return err
		}
	}

	if err := q.UpsertLocation(ctx, db.UpsertLocationParams{
		ID:          "foundation",
		Name:        "SCP Foundation Site-19",
		Description: "The largest and most active Foundation containment site. Hundreds of SCPs are held here.",
	}); err != nil {
		return err
	}

	rooms := []db.UpsertRoomParams{
		{ID: "containment-wing", LocationID: "foundation", Name: "Containment Wing", ImagePath: "/images/foundation/containment-wing.jpg"},
		{ID: "corridor", LocationID: "foundation", Name: "Main Corridor", ImagePath: "/images/foundation/corridor.jpg"},
		{ID: "cafeteria", LocationID: "foundation", Name: "Cafeteria", ImagePath: "/images/foundation/cafeteria.jpg"},
		{ID: "control-room", LocationID: "foundation", Name: "Control Room", ImagePath: "/images/foundation/control-room.jpg"},
		{ID: "medical-bay", LocationID: "foundation", Name: "Medical Bay", ImagePath: "/images/foundation/medical-bay.jpg"},
		{ID: "armory", LocationID: "foundation", Name: "Armory", ImagePath: "/images/foundation/armory.jpg"},
		{ID: "ventilation-shaft", LocationID: "foundation", Name: "Ventilation Shaft", ImagePath: "/images/foundation/ventilation-shaft.jpg"},
		{ID: "exit-checkpoint", LocationID: "foundation", Name: "Exit Checkpoint", ImagePath: "/images/foundation/exit-checkpoint.jpg"},
	}

	for _, r := range rooms {
		if err := q.UpsertRoom(ctx, r); err != nil {
			return err
		}
	}

	return nil
}
