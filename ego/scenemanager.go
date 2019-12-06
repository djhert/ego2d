package ego

var (
	scenes map[string]scene

	ActiveScene scene
	nextScene   scene

	sceneStarted bool
)

func init() {
	scenes = make(map[string]scene)
}

func AddScene(name string, s scene) {
	scenes[name] = s
}

func GetScene(name string) (scene, bool) {
	i, b := scenes[name]
	return i, b
}

func SetInitScene(name string) bool {
	if !isRunning {
		i, b := scenes[name]
		if b {
			ActiveScene = i
		}
		return b
	}
	return false
}

func SetNextScene(name string) bool {
	i, b := scenes[name]
	if b {
		nextScene = i
		return nextScene.Start()
	}
	return b
}

func ChangeScene(name string) bool {
	i, b := scenes[name]
	if b {
		ActiveScene = i
		sceneStarted = ActiveScene.Start()
	}
	return b
}

func NextScene() {
	sceneStarted = false
	ActiveScene = nextScene
	nextScene = nil
	sceneStarted = true
}
