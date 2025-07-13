import { User } from '../../domain/entities/User';
import type { HttpClient } from '../../domain/gateway/HttpClient';
import type { StorageClient } from '../../domain/gateway/StorageClient';
import { Email } from '../../domain/valuesobject/Email';
import { Password } from '../../domain/valuesobject/Password';
import type { RequestLoginUser } from '../dto/RequestLoginUserDTO';
import type { ResponseLoginUserDTO } from '../dto/ResponseLoginUserDTO';

export class LoginUserUseCase {
  public constructor(private httpClient: HttpClient, private storageClient: StorageClient) {}

  public async login(payload: RequestLoginUser): Promise<ResponseLoginUserDTO> {
    const email: Email = Email.create(payload.email);
    const password: Password = Password.create(payload.password);
    const user: User = User.create(email, password);
    const authenticationResponse = await this.httpClient.post<ResponseLoginUserDTO>(
      'login',
      user.getCredentialsForAuthentication()
    );
    this.storageClient.set('token', authenticationResponse.token);
    return authenticationResponse;
  }
}
