import { OnInit, Component } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

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
    //console.log(this.form.getRawValue());
    this.httpClient.post('/backend/users/register', this.form.getRawValue()).subscribe(() => {
      this.router.navigate(['/login'])});
  }
}
