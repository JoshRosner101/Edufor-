import { OnInit, Component } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { lastValueFrom } from 'rxjs'

//the variable names are the same ones in the databse
interface IThreadItem {
  id: number,
  username: string,
  title: string,
  body: string,
  time: string
}

@Component({
  selector: 'app-main-page',
  templateUrl: './main-page.component.html',
  styleUrls: ['./main-page.component.css']
})
export class MainPageComponent implements OnInit {
  public username = ''
  public title = ''
  public body = ''

  public threadItems: IThreadItem[] = []

  constructor(
    private httpClient: HttpClient
  ) {}

  async ngOnInit() {
    await this.loadNewItems()
  }

  async loadNewItems() {
    this.threadItems = await lastValueFrom(this.httpClient.get<IThreadItem[]>('/backend/threads'))
  }

  async addThread() {
    await lastValueFrom(this.httpClient.post('/backend/threads', {
      username: this.username,
      title: this.title,
      body: this.body,
      time: new Date().toLocaleString("en-US")
    }))

    await this.loadNewItems()
    
    this.username = ''
    this.title = ''
    this.body = ''
    
  }
}
