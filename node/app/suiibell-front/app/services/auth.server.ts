import { Authenticator, AuthorizationError } from 'remix-auth';
import { FormStrategy } from 'remix-auth-form';
import { sessionStorage, User } from '~/services/session.server';

const authenticator = new Authenticator<User | Error | null>(sessionStorage, {
  sessionKey: "sessionKey", // keep in sync
  sessionErrorKey: "sessionErrorKey", // keep in sync
});

authenticator.use(
  new FormStrategy(async ({ form }) => {

    let email = form.get('email') as string;
    let password = form.get('password') as string;

    let user = null;

    if (!email || email?.length === 0) throw new AuthorizationError('Bad Credentials: Email is required')
    if (typeof email !== 'string')
      throw new AuthorizationError('Bad Credentials: Email must be a string')

    if (!password || password?.length === 0) throw new AuthorizationError('Bad Credentials: Password is required')
    if (typeof password !== 'string')
      throw new AuthorizationError('Bad Credentials: Password must be a string')

    if (email === 'aaron@mail.com' && password === 'password') {
      user = {
        name: email,
        token: `${password}-${new Date().getTime()}`,
      };

      // the type of this user must match the type you pass to the Authenticator
      // the strategy will automatically inherit the type if you instantiate
      // directly inside the `use` method
      return await Promise.resolve({ ...user });

    } else {
      // if problem with user throw error AuthorizationError
      throw new AuthorizationError("Bad Credentials")
    }

  }),
);

export default authenticator