import { useFetcher } from 'react-router-dom';

import classes from './NewsletterSignup.module.css';
import { useEffect } from 'react';

function NewsletterSignup() {
  // MainNavigationの一部で全画面に関わるためuseFetcherで処理する
  // これによってRoot変更がなくなりSubmit後の繊維がなくなる

  // useFetcherは画面の更新なくSugmitしたい時に利用
  const fetcher = useFetcher();
  const { data, state } = fetcher;

  useEffect(() => {
    if (state === 'idle' && data && data.message) {
      window.alert(data.message);
    }
  }, [data, state]);

  // state === 'idle'

  return (
    <fetcher.Form method="post" action="/newsletter" className={classes.newsletter}>
      <input type="email" placeholder="Sign up for newsletter..." aria-label="Sign up for newsletter" />
      <button>Sign up</button>
    </fetcher.Form>
  );
}

export default NewsletterSignup;
