build:
	buildah bud -t hsmtkk/special-tribble

podman-play-kube:
	podman play kube pod.yml

podman-pod-rm:
	podman pod rm special-tribble -f
