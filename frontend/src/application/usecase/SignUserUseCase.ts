import { User } from '../../domain/entities/User';
import type { HttpPostClient } from '../../domain/gateway/HttpPostClient';
import { Email } from '../../domain/valuesobject/Email';
import { Name } from '../../domain/valuesobject/Name';
import { Password } from '../../domain/valuesobject/Password';
import type { RequestSignUserDTO } from '../dto/RequestSignUserDTO';
import type { ResponseSignUserDTO } from '../dto/ResponseSignUserDTO';

export class SignUserUseCase {
  public constructor(private httpClient: HttpPostClient, private path: string) {}

  public async sign(payload: RequestSignUserDTO): Promise<ResponseSignUserDTO> {
    const name: Name = Name.create(payload.name);
    const password: Password = Password.create(payload.password);
    const email: Email = Email.create(payload.email);
    const user: User = User.create(email, password, name);
    const registerResponse = await this.httpClient.post<ResponseSignUserDTO>(this.path, user.getCredentialsForRegister());
    return registerResponse;
  }
}
