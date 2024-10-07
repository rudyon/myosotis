<div class="hero-container">
	<img src="/api/placeholder/150/150" alt="Cute mascot" class="mascot" />

	<h1 class="hero-title">Welcome to Myosotis!</h1>
	<p class="hero-subtitle">Bloom your knowledge with our delightful flashcard experience ✨</p>

	<div class="features">
		<div class="feature-card">
			<h3>Adaptive Learning</h3>
			<p><add text about smart algorithm></add></p>
		</div>
		<div class="feature-card">
			<h3>Beautiful Design</h3>
			<p><add text about user experience></add></p>
		</div>
		<div class="feature-card">
			<h3>Study Together</h3>
			<p><add text about social features></add></p>
		</div>
	</div>

	<a href="/login" class="cta-button">Start Learning Today!</a>
</div>

<style>
	.hero-container {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		min-height: 90vh;
		text-align: center;
		padding: 2rem;
	}

	.hero-title {
		font-size: 3rem;
		margin-bottom: 0.5rem;
		background: linear-gradient(45deg, #a78bfa, #818cf8);
		-webkit-background-clip: text;
		background-clip: text;
		color: transparent;
		animation: fadeIn 1s ease-in;
	}

	.hero-subtitle {
		font-size: 1.25rem;
		margin-bottom: 2rem;
		color: #94a3b8;
		animation: fadeIn 1s ease-in 0.5s both;
	}

	.features {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
		gap: 2rem;
		margin-bottom: 2rem;
		animation: fadeIn 1s ease-in 1s both;
	}

	.feature-card {
		background: rgba(255, 255, 255, 0.05);
		border-radius: 1rem;
		padding: 1.5rem;
		backdrop-filter: blur(10px);
		transition: transform 0.3s;
	}

	.feature-card:hover {
		transform: translateY(-5px);
	}

	.cta-button {
		font-size: 1.25rem;
		padding: 0.75rem 2rem;
		border-radius: 9999px;
		background: linear-gradient(45deg, #a78bfa, #818cf8);
		color: white;
		text-decoration: none;
		transition:
			transform 0.3s,
			box-shadow 0.3s;
		animation: fadeIn 1s ease-in 1.5s both;
	}

	.cta-button:hover {
		transform: scale(1.05);
		box-shadow: 0 0 20px rgba(167, 139, 250, 0.5);
	}

	.mascot {
		width: 150px;
		height: 150px;
		margin-bottom: 1rem;
		animation: bounce 2s infinite;
	}

	@keyframes fadeIn {
		from {
			opacity: 0;
			transform: translateY(20px);
		}
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}

	@keyframes bounce {
		0%,
		100% {
			transform: translateY(0);
		}
		50% {
			transform: translateY(-10px);
		}
	}
</style>
