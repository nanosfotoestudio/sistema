import { useRouter } from 'next/router';
import useSWR from 'swr';
const fetcher = url => fetch(url).then(r => r.json());
export default function User({  }) {
  const router = useRouter();
  const { id } = router.query;
  const { data, error } = useSWR('/api/hello', fetcher)
  console.log(data);
  return <h1>Hello created at </h1>
}
