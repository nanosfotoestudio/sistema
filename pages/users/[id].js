import { useRouter } from 'next/router';

export default function User({ fecha }) {
  const router = useRouter();
  const { id } = router.query;

  return <h1>Hello {id} created at {fecha}</h1>
}

export async function getStaticProps() {
  return {
    props: { fecha: Date.now() }
  }
}

export async function getStaticPaths({ params }) {
  const ga = [1, 2, 3]
  const datos = ga.map( ga => {
    return { params: { id: ga } }
  })
  return {
    paths : ['/users/ga'],
    fallback: false
  };
}

