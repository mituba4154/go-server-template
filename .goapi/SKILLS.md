# Bukkit Java -> Go API (Go) Conversion Cheatsheet

Reference for converting common Bukkit Java patterns to the Go API used in this template.

## Availability Notice

GoAPI and GoLoader are not public yet.

---

## Event Handlers

| Bukkit Java | Go API (Go) |
|---|---|
| `@EventHandler void on(PlayerJoinEvent e)` | `goapi.On("PlayerJoinEvent", func(ctx *goapi.Context) {})` |
| `@EventHandler void on(PlayerQuitEvent e)` | `goapi.On("PlayerQuitEvent", func(ctx *goapi.Context) {})` |
| `@EventHandler void on(PlayerDeathEvent e)` | `goapi.On("PlayerDeathEvent", func(ctx *goapi.Context) {})` |
| `@EventHandler void on(PlayerMoveEvent e)` | `goapi.On("PlayerMoveEvent", func(ctx *goapi.Context) {})` |
| `@EventHandler void on(BlockBreakEvent e)` | `goapi.On("BlockBreakEvent", func(ctx *goapi.Context) {})` |
| `@EventHandler void on(EntityDeathEvent e)` | `goapi.On("EntityDeathEvent", func(ctx *goapi.Context) {})` |

Use the Bukkit class name as the event string, for example:
`PlayerInteractEvent` -> `"PlayerInteractEvent"`.

---

## Player Access

| Bukkit Java | Go API (Go) |
|---|---|
| `Player p = e.getPlayer()` | `p := ctx.Player()` |
| `Player p = e.getEntity()` (player event contexts) | `p := ctx.Player()` |

---

## Player Operations

| Bukkit Java | Go API (Go) |
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

## Command Registration

| Bukkit Java | Go API (Go) |
|---|---|
| `plugin.getCommand("home").setExecutor(...)` | `goapi.Command("home", func(ctx *goapi.CommandContext) {})` |
| `String[] args` | `ctx.Args()` |
| `CommandSender sender` | `ctx.Sender()` |
| `sender.sendMessage("usage")` | `ctx.Sender().SendMessage("usage")` |

---

## Scheduler

| Bukkit Java | Go API (Go) |
|---|---|
| `scheduler.runTaskLater(plugin, fn, 20L)` | `goapi.RunLater(20, func() { ... })` |
| `scheduler.runTaskTimer(plugin, fn, 0L, 20L)` | `goapi.RunTimer(0, 20, func() { ... })` |
| Run after 1 second | `goapi.RunLater(20, fn)` |
| Run every 5 minutes | `goapi.RunTimer(0, 6000, fn)` |

---

## World Operations

| Bukkit Java | Go API (Go) |
|---|---|
| `world.getBlockAt(x,y,z).getType().name()` | `goapi.World("world").GetBlock(x, y, z)` |
| `world.getBlockAt(x,y,z).setType(Material.STONE)` | `goapi.World("world").SetBlock(x, y, z, "STONE")` |
| `world.spawnEntity(loc, EntityType.ZOMBIE)` | `goapi.World("world").SpawnEntity(x, y, z, "ZOMBIE")` |
| `world.getSpawnLocation()` | `goapi.World("world").GetSpawn()` |
| `world.setSpawnLocation(x,y,z)` | `goapi.World("world").SetSpawn(x, y, z)` |
| `world.getTime()` | `goapi.World("world").GetTime()` |
| `world.setTime(6000L)` | `goapi.World("world").SetTime(6000)` |

---

## Material / EntityType Names

Pass Bukkit constant names as strings:

```go
// Material
p.GiveItem("DIAMOND_SWORD", 1)
p.GiveItem("IRON_INGOT", 64)
w.SetBlock(x, y, z, "OAK_LOG")

// EntityType
w.SpawnEntity(x, y, z, "ZOMBIE")
w.SpawnEntity(x, y, z, "CREEPER")
w.SpawnEntity(x, y, z, "ARMOR_STAND")
```
