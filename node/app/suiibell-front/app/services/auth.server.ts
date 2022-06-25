import { Authenticator, AuthorizationError } from 'remix-auth';
import { FormStrategy } from 'remix-auth-form';
import configuration from '~/conf/conf';
import { sessionStorage, User } from '~/services/session.server';

const authenticator = new Authenticator<User | Error | null>(sessionStorage, {
  sessionKey: "sessionKey", // keep in sync
  sessionErrorKey: "sessionErrorKey", // keep in sync
});

authenticator.use(
  new FormStrategy(async ({ form }) => {

    let email = form.get('email') as string;
    let password = form.get('password') as string;

    if (!email || email?.length === 0) throw new AuthorizationError('Bad Credentials: email is required')
    if (typeof email !== 'string')
      throw new AuthorizationError('Bad Credentials: user ID must be a string')

    if (!password || password?.length === 0) throw new AuthorizationError('Bad Credentials: Password is required')
    if (typeof password !== 'string')
      throw new AuthorizationError('Bad Credentials: Password must be a string')

    if (email === 'hogehoge@mail.com' && password === 'hogehoge1') {
      const hostIP = configuration['go-host']
      const hostPort = configuration['go-port']
      const hostPath = configuration['go-path']


      const userDataJson = await fetch(`${hostIP}:${hostPort}${hostPath}/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email: email,
          password: password,
        }),
      }).then(res => {
        if (res.status === 200) {
          return res.json();
        } else {
          throw new AuthorizationError('Bad Credentials: Invalid email or password');
        }
      }

      ).catch(err => {
        throw new AuthorizationError('Bad Credentials: Invalid email or password');
      }
      );

      const userData: User = {
        email: userDataJson.email,
        token: userDataJson.token,
      };



      // user = {
      //     email: email,
      //     token: `${password}-${new Date().getTime()}`,
      //   };

      // // the type of this user must match the type you pass to the Authenticator
      // // the strategy will automatically inherit the type if you instantiate
      // // directly inside the `use` method
      return await Promise.resolve({ ...userData});

    } else {
      // if problem with user throw error AuthorizationError
      throw new AuthorizationError("Bad Credentials")
    }

  }),
);

export default authenticator