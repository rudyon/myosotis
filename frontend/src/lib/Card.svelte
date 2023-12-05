<script>
	import { Get } from "../../wailsjs/go/main/Cards.js";

	let flipped = false;
	let resultText = "this is stupid";

	function flip() {
		flipped = !flipped;
	}

	function get() {
		flipped = false;
		Get("asshole").then((result) => (resultText = result));
	}
</script>

<button on:click={get}>function</button>

<div class="center">
	<div class="card-container" on:click={flip}>
		<div class="card front {flipped ? 'flipped_front' : 'front'}">
			<p>The pioneer of the experimental science of memory was...</p>
		</div>
		<div class="card back {flipped ? 'flipped_back' : 'back'}">
			<p>{resultText}</p>
		</div>
	</div>
</div>

<style>
	.center {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100vh;
	}

	.card-container {
		perspective: 1000px;
		width: 440px;
		height: 300px;
		position: relative;
		text-align: center;
	}

	.card {
		position: absolute;
		width: 100%;
		height: 100%;
		display: flex;
		justify-content: center;
		align-items: center;
		transition: transform 250ms;
		background-color: white;
		border-radius: 20px;
		color: black;
		font-size: 32px;
		backface-visibility: hidden;
	}

	.front {
		transform: rotateY(0deg);
	}

	.back {
		transform: rotateY(-180deg);
	}

	.flipped_back {
		transform: rotateY(0deg);
	}

	.flipped_front {
		transform: rotateY(180deg);
	}
</style>
