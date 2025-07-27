import { User } from '../../domain/entities/User';
import type { HttpClient } from '../../domain/gateway/HttpClient';
import { Email } from '../../domain/valuesobject/Email';
import { Name } from '../../domain/valuesobject/Name';
import { Password } from '../../domain/valuesobject/Password';
import type { RequestSignUserDTO } from '../dto/RequestSignUserDTO';
import type { ResponseSignUserDTO } from '../dto/ResponseSignUserDTO';

export class SignUserUseCase {
  public constructor(private httpClient: HttpClient) {}

  public async sign(payload: RequestSignUserDTO): Promise<ResponseSignUserDTO> {
    const name: Name = Name.create(payload.name);
    const password: Password = Password.create(payload.password);
    const email: Email = Email.create(payload.email);
    const user: User = User.create(email, password, name);
    const registerResponse = await this.httpClient.post<ResponseSignUserDTO>(payload.path, user.getCredentialsForRegister());
    return registerResponse;
  }
}
