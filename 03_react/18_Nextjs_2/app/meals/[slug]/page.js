export default function MelasDetailPage({ params }) {
  return (
    <main>
      <h1>Meals</h1>
      <p>{params.slug}</p>
    </main>
  );
}
