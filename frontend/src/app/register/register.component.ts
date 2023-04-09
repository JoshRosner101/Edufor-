import { OnInit, Component } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
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
      //Usernames must have more than one character
      name: ['', [Validators.required, Validators.pattern('\\S{1,}')]],
      //Passwords must have more than eight characters
      password: ['', [Validators.required, Validators.pattern('\\S{8,}')]]
    });
  }


  submit(): void {
    if(this.form.valid)
    {
      this.httpClient.post('/backend/users/register', this.form.getRawValue()).subscribe(response => {
        this.router.navigate(['/login']);
        console.log("gaming");
      }, error => {
        this.showError = true;
        if(error.status === 400)
        {
          this.message = "Username already exists";
        }
      }
      );
    }
  }
}
