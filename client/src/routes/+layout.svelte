<main>
	<slot />
</main>

<svelte:head>
	<style>
		@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap');

		:root {
			--bg-color: #0a0a0a;
			--text-color: #f5f5f5;
			--accent-color: #818cf8;
			--accent-color-2: #a78bfa;
			--card-bg: #1a1a1a;
			--border-color: #333;
		}

		* {
			box-sizing: border-box;
		}

		body {
			background-color: var(--bg-color);
			color: var(--text-color);
			font-family: 'Inter', sans-serif;
			line-height: 1.6;
			margin: 0;
			padding: 0;
		}

		h1,
		h2,
		h3,
		h4,
		h5,
		h6 {
			font-weight: 600;
			line-height: 1.2;
		}

		.card {
			background-color: var(--card-bg);
			border: 1px solid var(--border-color);
			border-radius: 8px;
			padding: 1.5rem;
			color: var(--text-color);
		}

		a {
			color: var(--accent-color);
			text-decoration: none;
			transition: color 0.3s;
		}

		a:hover {
			color: var(--accent-color-2);
		}

		button,
		.button {
			padding: 0.75rem 1.5rem;
			border: 2px solid var(--border-color);
			border-radius: 8px;
			font-size: 1rem;
			font-family: 'Inter', sans-serif;
			font-weight: 500;
			text-decoration: none;
			cursor: pointer;
			background-color: var(--card-bg);
			color: var(--text-color);
			transition: all 0.3s ease;
		}

		button:hover,
		.button:hover {
			background-color: var(--border-color);
			border-color: var(--accent-color);
			transform: translateY(-2px);
		}

		button:active,
		.button:active {
			background-color: var(--bg-color);
			transform: translateY(0);
		}

		button:disabled,
		.button:disabled {
			background-color: var(--card-bg);
			color: #666;
			cursor: not-allowed;
			border-color: var(--border-color);
			transform: none;
		}

		/* Utility classes */
		.container {
			max-width: 1200px;
			margin: 0 auto;
			padding: 0 1rem;
		}

		.gradient-text {
			background: linear-gradient(45deg, var(--accent-color), var(--accent-color-2));
			-webkit-background-clip: text;
			background-clip: text;
			color: transparent;
		}

		/* Responsive design */
		@media (max-width: 768px) {
			button,
			.button {
				width: 100%;
				text-align: center;
			}

			.container {
				padding: 0 0.5rem;
			}
		}
	</style>
</svelte:head>
