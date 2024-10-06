<main>
	<slot />
</main>

<svelte:head>
	<style>
		:root {
			--bg-color: #000;
			--text-color: #fff;
		}

		body {
			background-color: var(--bg-color);
			color: var(--text-color);
		}

		.card {
			color: #000;
		}

		a {
			color: #0ff;
		}

		button {
			padding: 10px 20px;
			border: 2px solid #444;
			border-radius: 5px;
			font-size: 16px;
			font-family: 'Arial', sans-serif;
			cursor: pointer;
			background-color: #333;
			color: #f5f5f5;
			transition:
				background-color 0.3s,
				color 0.3s,
				border-color 0.3s;
		}

		button:hover {
			background-color: #444;
			color: #ffffff;
			border-color: #666;
		}

		button:active {
			background-color: #222;
			border-color: #555;
		}

		button:disabled {
			background-color: #555;
			color: #888;
			cursor: not-allowed;
			border-color: #666;
		}
	</style>
</svelte:head>
