use std::io::{Read, Result, Write};

pub struct ReadStats<R> {
    reads: usize,
    bytes_through: usize,
    internal_buffer: R,
}

impl<R: Read> ReadStats<R> {
    pub fn new(internal_buffer: R) -> ReadStats<R> {
        ReadStats{
            reads: 0,
            bytes_through: 0,
            internal_buffer
        }
    }
    
    pub fn get_ref(&self) -> &R {
        &self.internal_buffer
    }

    pub fn bytes_through(&self) -> usize {
        self.bytes_through
    }

    pub fn reads(&self) -> usize {
        return self.reads
    }
}

impl<R: Read> Read for ReadStats<R> {
    fn read(&mut self, buf: &mut [u8]) -> Result<usize> {
        let result = self.internal_buffer.read(buf);
        if let Ok(bytes) = result {
            self.reads += 1;
            self.bytes_through += bytes;
        }
        result
    }
}

pub struct WriteStats<W> {
    bytes_through: usize,
    writes: usize,
    receiver: W,
}

impl<W: Write> WriteStats<W> {
    // _wrapped is ignored because W is not bounded on Debug or Display and therefore
    // can't be passed through format!(). For actual implementation you will likely
    // wish to remove the leading underscore so the variable is not ignored.
    pub fn new(receiver: W) -> WriteStats<W> {
        WriteStats{
            bytes_through: 0,
            writes: 0,
            receiver
        }
    }

    pub fn get_ref(&self) -> &W {
        &self.receiver
    }

    pub fn bytes_through(&self) -> usize {
        self.bytes_through
    }

    pub fn writes(&self) -> usize {
        self.writes
    }
}

impl<W: Write> Write for WriteStats<W> {
    fn write(&mut self, buf: &[u8]) -> Result<usize> {
        let result = self.receiver.write(&buf);
        if let Ok(written) = result {
            self.writes += 1;
            self.bytes_through += written;
        }
        result
    }

    fn flush(&mut self) -> Result<()> {
        self.receiver.flush()
    }
}
