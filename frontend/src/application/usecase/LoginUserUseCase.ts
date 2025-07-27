import { User } from '../../domain/entities/User';
import type { HttpPostClient } from '../../domain/gateway/HttpPostClient';
import type { StorageClient } from '../../domain/gateway/StorageClient';
import { Email } from '../../domain/valuesobject/Email';
import { Password } from '../../domain/valuesobject/Password';
import type { RequestLoginUserDTO } from '../dto/RequestLoginUserDTO';
import type { ResponseLoginUserDTO } from '../dto/ResponseLoginUserDTO';

export class LoginUserUseCase {
  public constructor(
    private httpClient: HttpPostClient,
    private storageClient: StorageClient,
    private path: string,
    private key: string
  ) {}

  public async login(payload: RequestLoginUserDTO): Promise<ResponseLoginUserDTO> {
    const email: Email = Email.create(payload.email);
    const password: Password = Password.create(payload.password);
    const user: User = User.create(email, password);
    const authenticationResponse = await this.httpClient.post<ResponseLoginUserDTO>(
      this.path,
      user.getCredentialsForAuthentication()
    );
    this.storageClient.set(this.key, { name: authenticationResponse.name, token: authenticationResponse.token });
    return authenticationResponse;
  }
}
