// This file is automatically generated by gen_packetIDs.go. DO NOT EDIT.

package data

//go:generate go run gen_packetIDs.go

// PktID represents a packet ID used in the minecraft protocol.
type PktID int32

// Valid PktID values.
const (
	// Clientbound packets for connections in the login state.
	EncryptionBeginClientbound PktID = 0x01
	Success                    PktID = 0x02
	Compress                   PktID = 0x03
	LoginPluginRequest         PktID = 0x04
	Disconnect                 PktID = 0x00
	// Serverbound packets for connections in the login state.
	LoginStart                 PktID = 0x00
	EncryptionBeginServerbound PktID = 0x01
	LoginPluginResponse        PktID = 0x02

	// Clientbound packets for connections in the play state.
	KickDisconnect             PktID = 0x19
	UnlockRecipes              PktID = 0x35
	Animation                  PktID = 0x05
	WorldEvent                 PktID = 0x21
	ScoreboardDisplayObjective PktID = 0x43
	CustomPayloadClientbound   PktID = 0x17
	NamedSoundEffect           PktID = 0x18
	HeldItemSlotClientbound    PktID = 0x3f
	ChatClientbound            PktID = 0x0e
	AbilitiesClientbound       PktID = 0x30
	EntityDestroy              PktID = 0x36
	SetPassengers              PktID = 0x4b
	KeepAliveClientbound       PktID = 0x1f
	SpawnEntityExperienceOrb   PktID = 0x01
	OpenHorseWindow            PktID = 0x1e
	RemoveEntityEffect         PktID = 0x37
	RelEntityMove              PktID = 0x27
	SelectAdvancementTab       PktID = 0x3c
	OpenSignEntity             PktID = 0x2e
	Map                        PktID = 0x25
	FacePlayer                 PktID = 0x33
	EntityEquipment            PktID = 0x47
	ResourcePackSend           PktID = 0x38
	NbtQueryResponse           PktID = 0x54
	ScoreboardObjective        PktID = 0x4a
	StopSound                  PktID = 0x52
	OpenWindow                 PktID = 0x2d
	Camera                     PktID = 0x3e
	Advancements               PktID = 0x57
	UpdateTime                 PktID = 0x4e
	Login                      PktID = 0x24
	PositionClientbound        PktID = 0x34
	UpdateViewPosition         PktID = 0x40
	EntitySoundEffect          PktID = 0x50
	Respawn                    PktID = 0x39
	BlockChange                PktID = 0x0b
	BlockBreakAnimation        PktID = 0x08
	Title                      PktID = 0x4f
	EntityTeleport             PktID = 0x56
	EntityEffect               PktID = 0x59
	TileEntityData             PktID = 0x09
	SpawnPosition              PktID = 0x42
	WorldBorder                PktID = 0x3d
	Experience                 PktID = 0x48
	PlayerlistHeader           PktID = 0x53
	WindowItems                PktID = 0x13
	EntityUpdateAttributes     PktID = 0x58
	EntityHeadRotation         PktID = 0x3a
	VehicleMoveClientbound     PktID = 0x2b
	MapChunk                   PktID = 0x20
	EntityLook                 PktID = 0x29
	Teams                      PktID = 0x4c
	UpdateViewDistance         PktID = 0x41
	Explosion                  PktID = 0x1b
	MultiBlockChange           PktID = 0x3b
	PlayerInfo                 PktID = 0x32
	CraftRecipeResponse        PktID = 0x2f
	TransactionClientbound     PktID = 0x11
	TradeList                  PktID = 0x26
	CloseWindowClientbound     PktID = 0x12
	TabCompleteClientbound     PktID = 0x0f
	SetCooldown                PktID = 0x16
	BlockAction                PktID = 0x0a
	NamedEntitySpawn           PktID = 0x04
	SpawnEntityPainting        PktID = 0x03
	UpdateLight                PktID = 0x23
	CombatEvent                PktID = 0x31
	SpawnEntityLiving          PktID = 0x02
	ScoreboardScore            PktID = 0x4d
	DeclareCommands            PktID = 0x10
	UpdateHealth               PktID = 0x49
	EntityMetadata             PktID = 0x44
	AttachEntity               PktID = 0x45
	Tags                       PktID = 0x5b
	EntityStatus               PktID = 0x1a
	AcknowledgePlayerDigging   PktID = 0x07
	Collect                    PktID = 0x55
	WorldParticles             PktID = 0x22
	Entity                     PktID = 0x2a
	UnloadChunk                PktID = 0x1c
	Difficulty                 PktID = 0x0d
	CraftProgressBar           PktID = 0x14
	BossBar                    PktID = 0x0c
	DeclareRecipes             PktID = 0x5a
	GameStateChange            PktID = 0x1d
	Statistics                 PktID = 0x06
	EntityVelocity             PktID = 0x46
	SetSlot                    PktID = 0x15
	OpenBook                   PktID = 0x2c
	SoundEffect                PktID = 0x51
	EntityMoveLook             PktID = 0x28
	SpawnEntity                PktID = 0x00
	// Serverbound packets for connections in the play state.
	EnchantItem                PktID = 0x08
	CustomPayloadServerbound   PktID = 0x0b
	SelectTrade                PktID = 0x23
	SetCreativeSlot            PktID = 0x28
	UpdateSign                 PktID = 0x2b
	WindowClick                PktID = 0x09
	PositionLook               PktID = 0x13
	UpdateCommandBlock         PktID = 0x26
	QueryBlockNbt              PktID = 0x01
	Flying                     PktID = 0x15
	KeepAliveServerbound       PktID = 0x10
	ClientCommand              PktID = 0x04
	BlockPlace                 PktID = 0x2e
	EntityAction               PktID = 0x1c
	PositionServerbound        PktID = 0x12
	ResourcePackReceive        PktID = 0x21
	Spectate                   PktID = 0x2d
	TeleportConfirm            PktID = 0x00
	GenerateStructure          PktID = 0x0f
	SetDifficulty              PktID = 0x02
	CloseWindowServerbound     PktID = 0x0a
	Look                       PktID = 0x14
	AdvancementTab             PktID = 0x22
	SetBeaconEffect            PktID = 0x24
	AbilitiesServerbound       PktID = 0x1a
	ChatServerbound            PktID = 0x03
	DisplayedRecipe            PktID = 0x1e
	RecipeBook                 PktID = 0x1f
	UpdateJigsawBlock          PktID = 0x29
	TransactionServerbound     PktID = 0x07
	SteerVehicle               PktID = 0x1d
	NameItem                   PktID = 0x20
	PickItem                   PktID = 0x18
	UpdateStructureBlock       PktID = 0x2a
	TabCompleteServerbound     PktID = 0x06
	HeldItemSlotServerbound    PktID = 0x25
	SteerBoat                  PktID = 0x17
	Settings                   PktID = 0x05
	UseItem                    PktID = 0x2f
	CraftRecipeRequest         PktID = 0x19
	UpdateCommandBlockMinecart PktID = 0x27
	BlockDig                   PktID = 0x1b
	EditBook                   PktID = 0x0c
	UseEntity                  PktID = 0x0e
	VehicleMoveServerbound     PktID = 0x16
	ArmAnimation               PktID = 0x2c
	LockDifficulty             PktID = 0x11
	QueryEntityNbt             PktID = 0x0d

	// Clientbound packets used to respond to ping/status requests.
	ServerInfo      PktID = 0x00
	PingClientbound PktID = 0x01
	// Serverbound packets used to ping or read server status.
	PingStart       PktID = 0x00
	PingServerbound PktID = 0x01
)
