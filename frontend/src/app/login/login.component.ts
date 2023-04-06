import { OnInit, Component } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { Emitters } from '../emitters/emitters';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit{
  form!: FormGroup;


  constructor(
    private httpClient: HttpClient,
    private formBuilder: FormBuilder,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.form = this.formBuilder.group({
      name: '',
      password: ''
    });
  }


  submit(): void {
    this.httpClient.post('/backend/users/login', this.form.getRawValue(), {withCredentials: true}).subscribe(() => {
      this.httpClient.get('/backend/users/user', {withCredentials: true}).subscribe(
        (res: any) => {
          Emitters.authEmitter.emit(true);
        },
        err => {
          Emitters.authEmitter.emit(false);
        }
      );
      this.router.navigate(['/home'])});
  }
}
