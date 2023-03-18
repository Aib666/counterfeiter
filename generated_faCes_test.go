$ cat tools/tools.go

// +build tools

package tools

import (
	_ "github.com/maxbrunsfeld/counterfeiter/v6"
)

// This file imports packages that are used when running go generate, or used
// during the development process but not otherwise depended on by built code.

$ cat myinterface.go

package foo

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . MySpecialInterface

type MySpecialInterface interface {
	DoThings(string, uint64) (int, error)
}

$ go generate ./...
Writing `FakeMySpecialInterface` to `foofakes/fake_my_special_interface.go`... Done

$ cat myinterface.go

	package foo

// You only need **one** of these per package!
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// You will add lots of directives like these in the same package...
//counterfeiter:generate . MySpecialInterface
type MySpecialInterface interface {
	DoThings(string, uint64) (int, error)
}

// Like this...
//counterfeiter:generate . MyOtherInterface
type MyOtherInterface interface {
	DoOtherThings(string, uint64) (int, error)
}

$ go generate ./...
Writing `FakeMySpecialInterface` to `foofakes/fake_my_special_interface.go`... Done
Writing `FakeMyOtherInterface` to `foofakes/fake_my_other_interface.go`... Done

$ go generate ./...
$ go run github.com/maxbrunsfeld/counterfeiter/v6

USAGE
	counterfeiter
		[-generate] [-o <output-path>] [-p] [--fake-name <fake-name>]
		[<source-path>] <interface> [-]

$ GO111MODULE=off go get -u github.com/maxbrunsfeld/counterfeiter
$ counterfeiter

USAGE
	counterfeiter
		[-generate] [-o <output-path>] [-p] [--fake-name <fake-name>]
		[<source-path>] <interface> [-]

$ cat path/to/foo/file.go
package foo

type MySpecialInterface interface {
		DoThings(string, uint64) (int, error)
}

$ go run github.com/maxbrunsfeld/counterfeiter/v6 path/to/foo MySpecialInterface
Wrote `FakeMySpecialInterface` to `path/to/foo/foofakes/fake_my_special_interface.go`

import "my-repo/path/to/foo/foofakes"

var fake = &foofakes.FakeMySpecialInterface{}

fake.DoThings("stuff", 5)

Expect(fake.DoThingsCallCount()).To(Equal(1))

str, num := fake.DoThingsArgsForCall(0)
Expect(str).To(Equal("stuff"))
Expect(num).To(Equal(uint64(5)))

fake.DoThingsReturns(3, errors.New("the-error"))

num, err := fake.DoThings("stuff", 5)
Expect(num).To(Equal(3))
Expect(err).To(Equal(errors.New("the-error")))

	it("records a nil slice argument as a nil", func() {
		var buffer []byte = nil

		fake.DoASlice(buffer)

		arg1 := fake.DoASliceArgsForCall(0)
		Expect(arg1).To(BeNil())
	})

	it("records an array argument as a copy", func() {
		buffer := [4]byte{1, 2, 3, 4}

		fake.DoAnArray(buffer)

		buffer[0] = 2
		arg1 := fake.DoAnArrayArgsForCall(0)
		Expect(arg1).To(ConsistOf(byte(1), byte(2), byte(3), byte(4)))
	})

	it("passes the original slice to a stub function", func() {
		buffer := []byte{1}

		fake.DoASliceStub = func(b []byte) {
			b[0] = 2
		}

		fake.DoASlice(buffer)

		arg1 := fake.DoASliceArgsForCall(0)
		Expect(arg1).To(ConsistOf(byte(1)))
		Expect(buffer).To(ConsistOf(byte(2)))
	})

	it("records its calls without race conditions", func() {
		go fake.DoNothing()

		Eventually(fake.DoNothingCallCount, 1.0).Should(Equal(1))
	})

	when("implementing an interface to show recorded methoded invocations", func() {
		it.Before(func() {
			var ifake interface{} = fake
			_, ok := ifake.(InvocationRecorder)

			Expect(ok).To(BeTrue())
		})

		it("records each invocation", func() {
			Expect(len(fake.Invocations()["DoThings"])).To(Equal(0))
			Expect(len(fake.Invocations()["DoNothing"])).To(Equal(0))
			Expect(len(fake.Invocations()["DoASlice"])).To(Equal(0))
			Expect(len(fake.Invocations()["DoAnArray"])).To(Equal(0))

			fake.DoThings("hello", 0)
			Expect(len(fake.Invocations()["DoThings"])).To(Equal(1))
			Expect(fake.Invocations()["DoThings"][0][0]).To(Equal("hello"))
			Expect(fake.Invocations()["DoThings"][0][1]).To(Equal(uint64(0)))

			fake.DoNothing()
			Expect(len(fake.Invocations()["DoNothing"])).To(Equal(1))

			fake.DoASlice([]byte("HAI"))
			Expect(len(fake.Invocations()["DoASlice"])).To(Equal(1))
			Expect(fake.Invocations()["DoASlice"][0][0]).To(Equal([]byte("HAI")))

			fake.DoAnArray([4]byte{})
			Expect(len(fake.Invocations()["DoAnArray"])).To(Equal(1))
			Expect(fake.Invocations()["DoAnArray"][0][0]).ToNot(BeNil())
		})
	})

	when("when two methods are called at the same time", func() {
		var start1 chan struct{}
		var start2 chan struct{}
		var end1 chan struct{}
		var end2 chan struct{}

		it.Before(func() {
			start1 = make(chan struct{})
			end1 = make(chan struct{})
			start2 = make(chan struct{})
			end2 = make(chan struct{})

			fake.DoNothingStub = func() {
				close(start1)
				end1 <- struct{}{}
			}

			fake.DoThingsStub = func(string, uint64) (int, error) {
				close(start2)
				end2 <- struct{}{}
				return 0, nil
			}

			go fake.DoNothing()
			<-start1
			go fake.DoThings("abc", 1)
			<-start2
		})

		it.After(func() {
			close(end1)
			close(end2)
		})

		it("does not deadlock", func() {
			Eventually(start1).Should(BeClosed())
			Eventually(end1).Should(Receive())

			Eventually(start2).Should(BeClosed())
			Eventually(end2).Should(Receive())
		})
	})

	when("when methods are called concurrently", func() {
		it.Before(func() {
			go fake.DoNothing()
			go fake.DoThings("", 0)
		})

		it("records the call count without race conditions", func() {
			Eventually(fake.DoNothingCallCount).Should(Equal(1))
			Eventually(fake.DoThingsCallCount).Should(Equal(1))
		})

		it("records the invocations without race conditions as well", func() {
			Eventually(func() [][]interface{} { return fake.Invocations()["DoNothing"] }).Should(HaveLen(1))
			Eventually(func() [][]interface{} { return fake.Invocations()["DoThings"] }).Should(HaveLen(1))
		})
	})

	when("when the same method is called concurrently", func() {
		it("does not deadlock", func() {
			Eventually(func() bool {
				a := make(chan struct{})
				b := make(chan struct{})

				fake.DoNothingStub = func() {
					select {
					case <-a:
						close(b)
					default:
						close(a)
						<-b
					}
				}

				go fake.DoNothing()
				go fake.DoNothing()

				<-b
				return true
			}).Should(Equal(true))
		})
	})

	when("interfaces with var-args methods", func() {
		var fake *fixturesfakes.FakeHasVarArgs

		it.Before(func() {
			fake = new(fixturesfakes.FakeHasVarArgs)
		})

		it("implements the interface", func() {
			var interfaceVal fixtures.HasVarArgs = fake
			Expect(interfaceVal).NotTo(BeNil())
		})

		it("records the calls in a slice", func() {
			fake.DoThings(5, "one", "two", "three")

			num, strings := fake.DoThingsArgsForCall(0)
			Expect(num).To(Equal(5))
			Expect(strings).To(Equal([]string{"one", "two", "three"}))
		})

		it("passes the var-args to stub functions", func() {
			fake.DoThingsStub = func(x int, strings ...string) int {
				Expect(strings).To(Equal([]string{"one", "two", "three"}))
				return 11
			}

			val := fake.DoThings(5, "one", "two", "three")
			Expect(val).To(Equal(11))
		})
	})

	when("calling a function repeatedly and changing the returns value", func() {
		var fake *fixturesfakes.FakeSomething

		it.Before(func() {
			fake = new(fixturesfakes.FakeSomething)
		})

		it("succeeds and does not cause a data race", func() {
			go fake.DoThingsReturns(1, nil)
			go fake.DoThings("1", 1)
			go fake.DoThingsReturns(1, nil)
			go fake.DoThings("1", 1)
		})
	})
}

type InvocationRecorder interface {
	Invocations() map[string][][]interface{}
}
