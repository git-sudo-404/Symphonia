/*
 * MIT License
 *
 * Copyright (c) 2026 git-sudo-404
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * Of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * Copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * Copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING, BUT NOT LIMITED TO, THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES, OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package internal

type LogEntry struct {
	term uint64
	cmd  string
}

type Node struct {
	// PERSISTENT STATES _ on all servers
	// (updated on stable storage before responding to RPC's)
	currentTerm uint64 // the latest term a server has seen, initialized with 0 and increases monotonically
	votedFor    string // candidateId that recieved the vote of this server in the current term
	log         []LogEntry

	// VOLATILE STATES _ on all servers
	commitIndex uint64 // index of the highest known log entry to be committed by the quorom, initialized with 0 and increases monotonically
	lastApplied uint64 // index of the highest log entry applied to the state machine, initialized with 0 and increases monotonically

	// VOLATILE STATES _ only on leaders
	// (reinitialized after election)
	nextIndex  map[string]uint64 // for each peer, the index of the next entry to be sent to that peer, initialized to leader last log index + 1
	matchIndex map[string]uint64 // for each peer, the index of the highest log entry known to be replicated on the peer, initialized with 0 and increases monotonically
}
