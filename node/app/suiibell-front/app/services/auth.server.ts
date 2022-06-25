import { Authenticator, AuthorizationError } from 'remix-auth';
import { FormStrategy } from 'remix-auth-form';
import configuration from '~/conf/conf';
import { sessionStorage, User } from '~/services/session.server';

// authenticator is a singleton, so we can use it in multiple places
const authenticator = new Authenticator<User | Error | null>(sessionStorage, {
  sessionKey: "sessionKey", // keep in sync
  sessionErrorKey: "sessionErrorKey", // keep in sync
});

// "use" is the keyword for a function that is used to register a user data
authenticator.use(
  new FormStrategy(async ({ form }) => {

    let email = form.get('email') as string;
    let password = form.get('password') as string;

    if (!email || email?.length === 0) throw new AuthorizationError('Bad Credentials: email is required')
    if (typeof email !== 'string')
      throw new AuthorizationError('Bad Credentials: user ID must be a string')

    if (!password || password?.length === 0) throw new AuthorizationError('Bad Credentials: Password is required')
    if (typeof password !== 'string') {
      throw new AuthorizationError('Bad Credentials: Password must be a string')

    }

    const hostIP = configuration['go-host']
    const hostPort = configuration['go-port']
    const hostPath = configuration['go-path']

    // call API to get user data
    // Go is a backend server that runs on port 8000 and handles requests
    // to the /api/v1 path.
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
    });

    const userData: User = {
      email: userDataJson.email,
      token: userDataJson.token,
    };


    console.log(userData);

    return await Promise.resolve({ ...userData });
  })
);

export default authenticator