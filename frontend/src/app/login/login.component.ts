import { OnInit, Component } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Emitters } from '../emitters/emitters';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit{
  form!: FormGroup;
  showError: boolean = false;
  message: string = "";


  constructor(
    private httpClient: HttpClient,
    private formBuilder: FormBuilder,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.form = this.formBuilder.group({
      name: ['', Validators.required],
      password: ['', Validators.required]
    });
  }


  submit(): void {
    if(this.form.valid)
    {
      this.httpClient.post('/backend/users/login', this.form.getRawValue(), {withCredentials: true, observe: 'response'}).subscribe(response => {
        this.httpClient.get('/backend/users/user', {withCredentials: true}).subscribe(
          (res: any) => {
            Emitters.authEmitter.emit(true);
          },
          err => {
            Emitters.authEmitter.emit(false);
          }
        );
        this.router.navigate(['/home']);
      }, error => {
        this.showError = true;
        if(error.status === 400)
        {
          this.message = "Incorrect Password";
        }
        if(error.status === 404)
        {
          this.message = "Username does not exist";
        }
      });
    }
  }
}
