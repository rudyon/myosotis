<script>
	let isFlipped = false;

	function flipCard() {
		isFlipped = !isFlipped;
	}
</script>

<div id="cards">
	<div id="current_card">
		<div class="scale_on_hover" on:click={flipCard}>
			<div id="flip-container" flip={isFlipped ? 'yes' : 'no'}>
				<div class="flipper">
					<div id="card_front" class="card">
						<p>front</p>
					</div>
					<div id="card_back" class="card">
						<p>back</p>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>

<style>
	#cards {
		position: absolute;
		left: 100px;
		top: 35px;
	}

	#current_card {
		top: 0;
		width: 400px;
		height: 240px;
		transition: left 0.3s ease-in-out;
		position: relative;
		left: 0;
	}

	.card {
		width: 396px;
		height: 236px;
		border-radius: 20px;
		overflow: hidden;
		font-size: 40px;
		border: 2px inset #000;
		text-align: center;
		backface-visibility: hidden; /* Hide the back of each card when it's flipped */
		position: absolute; /* Stack front and back cards on top of each other */
		top: 0;
		left: 0;
		display: flex; /* Enable flexbox */
		justify-content: center; /* Center horizontally */
		align-items: center; /* Center vertically */
	}

	.scale_on_hover {
		width: 100%;
		height: 100%;
		transition: 0.15s;
		transform-origin: 200px 120px;
		cursor: pointer;
	}

	.scale_on_hover:hover {
		transform: scale(1.03);
	}

	#flip-container {
		width: 100%;
		height: 100%;
		perspective: 1200px;
		position: relative;
	}

	#flip-container[flip='yes'] .flipper {
		transform: rotateY(180deg);
	}

	#flip-container[flip='no'] .flipper {
		transform: rotateY(0deg);
	}

	#flip-container .flipper {
		width: 100%;
		height: 100%;
		transition: 0.5s;
		transform-style: preserve-3d;
		position: relative; /* Ensure the cards stay within the container */
	}

	#card_front {
		z-index: 2;
		transform: rotateY(0deg);
		background: #fff;
	}

	#card_back {
		transform: rotateY(180deg);
		background: #fff;
	}
</style>
