import { OnInit, Component } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { HttpClientTestingModule } from '@angular/common/http/testing'
import { lastValueFrom } from 'rxjs'
import { ActivatedRoute } from '@angular/router'

interface IReplyItem {
  replyid: number,
  username: string,
  body: string,
  time: string,
  replypost: number
}

//the variable names are the same ones in the databse
interface IThreadItem {
  id: number,
  username: string,
  title: string,
  body: string,
  time: string,
  replies: IReplyItem[]
}

@Component({
  selector: 'app-by-id',
  templateUrl: './by-id.component.html',
  styleUrls: ['./by-id.component.css']
})

export class ByIDComponent implements OnInit{
  public username = ''
  public title = ''
  public body = ''

  public threadItem: IThreadItem = {id: 0, username: "", title: "", body: "", time: "", replies: []}

  constructor(
    private httpClient: HttpClient,
    private route: ActivatedRoute
  ) {}

  async ngOnInit() {
    await this.loadNewItems()
  }

  async loadNewItems() {
    const urlpath = "/backend/threads/" + this.route.snapshot.paramMap.get('id')
    this.threadItem = await lastValueFrom(this.httpClient.get<IThreadItem>(urlpath))
  }

  async addReply() {
    let currentID = this.route.snapshot.paramMap.get('id')
    const urlpath = "/backend/threads/" + currentID
    await lastValueFrom(this.httpClient.post(urlpath, {
      username: this.username,
      body: this.body,
      time: new Date().toLocaleString("en-US"),
      replypost: Number(currentID)
    }))

    await this.loadNewItems()
    
    this.username = ''
    this.body = ''
  }
}
