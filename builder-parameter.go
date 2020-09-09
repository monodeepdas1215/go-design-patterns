package main

/*
There are cases when we do not want to expose out struct entity and only allow interaction with it through builders.
This is where the Builder parameter comes in.
*/

// this struct is not exposed outside of the package
type email struct {
	From, To, Heading, Body string
}

// this builder is exposed publicly
type EmailBuilder struct {
	email *email
}

func (eb *EmailBuilder) From(from string) *EmailBuilder {
	eb.email.From = from
	return eb
}

func (eb *EmailBuilder) To(to string) *EmailBuilder {
	eb.email.To = to
	return eb
}

func (eb *EmailBuilder) Heading(h string) *EmailBuilder {
	eb.email.Heading = h
	return eb
}

func (eb *EmailBuilder) Body(b string) *EmailBuilder {
	eb.email.Body = b
	return eb
}

func (eb *EmailBuilder) Email() *email {
	return eb.email
}


// THIS IS THE ACTUAL BUILDER PARAMETER
type EmailBuilderParam func(builder *EmailBuilder)


// This function accepts a function parameter of the email builder
func SendEmail(action EmailBuilderParam) {
	builder := EmailBuilder{}
	action(&builder)
	sendEmailImplemntation(builder.Email())
}

func sendEmailImplemntation(email *email) {
}

func main() {

	SendEmail(func(builder *EmailBuilder) {
		builder.From("monodeepdas112@gmail.com").To("abc@gmail.com").Heading("meeting").Body("lets meet !")
	})
}