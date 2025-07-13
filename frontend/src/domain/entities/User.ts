import type { Email } from '../valuesobject/Email';
import type { Name } from '../valuesobject/Name';
import type { Password } from '../valuesobject/Password';

export class User {
  private constructor(private email: Email, private password: Password, private name?: Name) {}

  public static create(email: Email, password: Password, name?: Name): User {
    return new User(email, password, name);
  }

  public getName(): string {
    if (!this.name) throw new Error('name undefined');
    return this.name.getValue();
  }

  public getCredentialsForAuthentication() {
    return {
      password: this.password.getValue(),
      email: this.email.getValue(),
    };
  }
  public getCredentialsForRegister() {
    return {
      name: this.getName(),
      password: this.password.getValue(),
      email: this.email.getValue(),
    };
  }
}
