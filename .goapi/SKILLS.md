# Bukkit Java → GoAPILib Go 変換チートシート

AI向け変換リファレンス。
Bukkit の Java コードを GoAPILib の Go コードに変換するときに参照すること。

---

## イベントハンドラ

| Bukkit Java | GoAPILib Go |
|---|---|
| `@EventHandler void on(PlayerJoinEvent e)` | `goapi.On("PlayerJoinEvent", func(ctx *goapi.Context) {})` |
| `@EventHandler void on(PlayerQuitEvent e)` | `goapi.On("PlayerQuitEvent", func(ctx *goapi.Context) {})` |
| `@EventHandler void on(PlayerDeathEvent e)` | `goapi.On("PlayerDeathEvent", func(ctx *goapi.Context) {})` |
| `@EventHandler void on(PlayerMoveEvent e)` | `goapi.On("PlayerMoveEvent", func(ctx *goapi.Context) {})` |
| `@EventHandler void on(BlockBreakEvent e)` | `goapi.On("BlockBreakEvent", func(ctx *goapi.Context) {})` |
| `@EventHandler void on(EntityDeathEvent e)` | `goapi.On("EntityDeathEvent", func(ctx *goapi.Context) {})` |

イベント名は Bukkit のクラス名をそのまま文字列にする。
例: `PlayerInteractEvent` -> `"PlayerInteractEvent"`

---

## プレイヤー取得

| Bukkit Java | GoAPILib Go |
|---|---|
| `Player p = e.getPlayer()` | `p := ctx.Player()` |
| `Player p = e.getEntity()` (PlayerEvent系) | `p := ctx.Player()` |

---

## プレイヤー操作

| Bukkit Java | GoAPILib Go |
|---|---|
| `p.sendMessage("hello")` | `p.SendMessage("hello")` |
| `p.getName()` | `p.Name()` |
| `p.getUniqueId().toString()` | `p.ID()` |
| `p.getHealth()` | `p.GetHealth()` |
| `p.setHealth(20.0)` | `p.SetHealth(20.0)` |
| `p.teleport(location)` | `p.Teleport("world", x, y, z)` |
| `p.getLocation()` | `p.GetLocation()` -> `Location{World, X, Y, Z}` |
| `p.kickPlayer("reason")` | `p.Kick("reason")` |
| `p.getInventory().addItem(new ItemStack(Material.STONE, 64))` | `p.GiveItem("STONE", 64)` |
| `p.setGameMode(GameMode.CREATIVE)` | `p.SetGameMode("CREATIVE")` |
| `p.getGameMode().name()` | `p.GetGameMode()` |

---

## コマンド登録

| Bukkit Java | GoAPILib Go |
|---|---|
| `plugin.getCommand("home").setExecutor(...)` | `goapi.Command("home", func(ctx *goapi.CommandContext) {})` |
| `String[] args` | `ctx.Args()` |
| `CommandSender sender` | `ctx.Sender()` |
| `sender.sendMessage("usage")` | `ctx.Sender().SendMessage("usage")` |

---

## スケジューラ

| Bukkit Java | GoAPILib Go |
|---|---|
| `scheduler.runTaskLater(plugin, fn, 20L)` | `goapi.RunLater(20, func() { ... })` |
| `scheduler.runTaskTimer(plugin, fn, 0L, 20L)` | `goapi.RunTimer(0, 20, func() { ... })` |
| 1秒後に実行 | `goapi.RunLater(20, fn)` |
| 5分ごとに実行 | `goapi.RunTimer(0, 6000, fn)` |

---

## ワールド操作

| Bukkit Java | GoAPILib Go |
|---|---|
| `world.getBlockAt(x,y,z).getType().name()` | `goapi.World("world").GetBlock(x, y, z)` |
| `world.getBlockAt(x,y,z).setType(Material.STONE)` | `goapi.World("world").SetBlock(x, y, z, "STONE")` |
| `world.spawnEntity(loc, EntityType.ZOMBIE)` | `goapi.World("world").SpawnEntity(x, y, z, "ZOMBIE")` |
| `world.getSpawnLocation()` | `goapi.World("world").GetSpawn()` |
| `world.setSpawnLocation(x,y,z)` | `goapi.World("world").SetSpawn(x, y, z)` |
| `world.getTime()` | `goapi.World("world").GetTime()` |
| `world.setTime(6000L)` | `goapi.World("world").SetTime(6000)` |

---

## Material / EntityType 名

Bukkit の定数名をそのまま文字列で渡す。

```go
// Material
p.GiveItem("DIAMOND_SWORD", 1)
p.GiveItem("IRON_INGOT", 64)
w.SetBlock(x, y, z, "OAK_LOG")

// EntityType
w.SpawnEntity(x, y, z, "ZOMBIE")
w.SpawnEntity(x, y, z, "CREEPER")
w.SpawnEntity(x, y, z, "ARMOR_STAND")
